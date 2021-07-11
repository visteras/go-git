package dotgit

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage"
	"github.com/go-git/go-git/v5/storage/s3/wrapper"
	"github.com/go-git/go-git/v5/utils/ioutil"
)

func (d *DotGit) SetRef(r, old *plumbing.Reference) error {
	var content string
	switch r.Type() {
	case plumbing.SymbolicReference:
		content = fmt.Sprintf("ref: %s\n", r.Target())
	case plumbing.HashReference:
		content = fmt.Sprintln(r.Hash().String())
	}

	fileName := r.Name().String()

	return d.setRef(fileName, content, old)
}

func (d *DotGit) setRef(fileName string, content string, old *plumbing.Reference) error {
	// If we are not checking an old ref, just truncate the file.
	if old == nil {
		err := d.s3.Delete(fileName)
		if err != nil {
			return err
		}
	}

	b, err := d.s3.Download(fileName)
	if err != nil {
		return err
	}

	// this is a no-op to call even when old is nil.
	err = d.checkReferenceAndTruncate(b, old)
	if err != nil {
		return err
	}

	return d.s3.Upload(fileName, []byte(content))
}

func (d *DotGit) checkReferenceAndTruncate(b []byte, old *plumbing.Reference) error {
	if old == nil {
		return nil
	}
	ref, err := d.readReferenceFrom(b, old.Name().String())
	if err != nil {
		return err
	}
	if ref.Hash() != old.Hash() {
		return storage.ErrReferenceHasChanged
	}
	return nil
}

func (d *DotGit) readReferenceFrom(b []byte, name string) (*plumbing.Reference, error) {
	line := strings.TrimSpace(string(b))
	return plumbing.NewReferenceFromStrings(name, line), nil
}

func (d *DotGit) Ref(name plumbing.ReferenceName) (*plumbing.Reference, error) {
	ref, err := d.readReferenceFile(".", name.String())
	if err == nil {
		return ref, nil
	}

	return d.packedRef(name)
}

func (d *DotGit) readReferenceFile(path, name string) (*plumbing.Reference, error) {
	path = d.s3.Join(path, d.s3.Join(strings.Split(name, "/")...))

	b, err := d.s3.Download(path)
	if err != nil {
		return nil, err
	}

	return d.readReferenceFrom(b, name)
}

func (d *DotGit) packedRef(name plumbing.ReferenceName) (*plumbing.Reference, error) {
	refs, err := d.findPackedRefs()
	if err != nil {
		return nil, err
	}

	for _, ref := range refs {
		if ref.Name() == name {
			return ref, nil
		}
	}

	return nil, plumbing.ErrReferenceNotFound
}

func (d *DotGit) findPackedRefs() (r []*plumbing.Reference, err error) {
	b, err := d.s3.Download(packedRefsPath)
	if err != nil {
		if errors.Is(err, wrapper.ErrNoSuchKey) {
			return nil, nil
		}
		return nil, err
	}

	return d.findPackedRefsInFile(b)
}

func (d *DotGit) findPackedRefsInFile(b []byte) ([]*plumbing.Reference, error) {
	buf := bytes.NewBuffer(b)
	sc := bufio.NewScanner(buf)

	var refs []*plumbing.Reference
	for sc.Scan() {
		ref, err := d.processLine(sc.Text())
		if err != nil {
			return nil, err
		}

		if ref != nil {
			refs = append(refs, ref)
		}
	}

	return refs, sc.Err()
}

func (d *DotGit) processLine(line string) (*plumbing.Reference, error) {
	if len(line) == 0 {
		return nil, nil
	}

	switch line[0] {
	case '#': // comment - ignore
		return nil, nil
	case '^': // annotated tag commit of the previous line - ignore
		return nil, nil
	default:
		ws := strings.Split(line, " ") // hash then ref
		if len(ws) != 2 {
			return nil, ErrPackedRefsBadFormat
		}

		return plumbing.NewReferenceFromStrings(ws[1], ws[0]), nil
	}
}

func (d *DotGit) Refs() ([]*plumbing.Reference, error) {
	var refs []*plumbing.Reference
	var seen = make(map[plumbing.ReferenceName]bool)
	if err := d.addRefsFromRefDir(&refs, seen); err != nil {
		return nil, err
	}

	if err := d.addRefsFromPackedRefs(&refs, seen); err != nil {
		return nil, err
	}

	if err := d.addRefFromHEAD(&refs); err != nil {
		return nil, err
	}

	return refs, nil
}

func (d *DotGit) addRefsFromRefDir(refs *[]*plumbing.Reference, seen map[plumbing.ReferenceName]bool) error {
	return d.walkReferencesTree(refs, []string{refsPath}, seen)
}

func (d *DotGit) walkReferencesTree(refs *[]*plumbing.Reference, relPath []string, seen map[plumbing.ReferenceName]bool) error {
	files, err := d.s3.GetListObjects(d.s3.Join(relPath...))
	if err != nil {
		if errors.Is(err, wrapper.ErrNoSuchKey) {
			return nil
		}

		return err
	}

	for _, f := range files {
		newRelPath := append(append([]string(nil), relPath...), f)

		ref, err := d.readReferenceFile(".", d.s3.Join(newRelPath...))
		if err != nil {
			return err
		}

		if ref != nil && !seen[ref.Name()] {
			*refs = append(*refs, ref)
			seen[ref.Name()] = true
		}
	}

	return nil
}

func (d *DotGit) addRefsFromPackedRefs(refs *[]*plumbing.Reference, seen map[plumbing.ReferenceName]bool) error {
	packedRefs, err := d.findPackedRefs()
	if err != nil {
		return err
	}

	for _, ref := range packedRefs {
		if !seen[ref.Name()] {
			*refs = append(*refs, ref)
			seen[ref.Name()] = true
		}
	}
	return nil
}

func (d *DotGit) addRefFromHEAD(refs *[]*plumbing.Reference) error {
	ref, err := d.readReferenceFile(".", "HEAD")
	if err != nil {
		if errors.Is(err, wrapper.ErrNoSuchKey) {
			return nil
		}

		return err
	}

	*refs = append(*refs, ref)
	return nil
}

func (d *DotGit) RemoveRef(name plumbing.ReferenceName) error {
	path := d.s3.Join(".", name.String())
	err := d.s3.Delete(path)
	if err != nil && !errors.Is(err, wrapper.ErrNoSuchKey) {
		return err
	}

	return d.rewritePackedRefsWithoutRef(name)
}

func (d *DotGit) rewritePackedRefsWithoutRef(name plumbing.ReferenceName) error {
	pr, err := d.openAndLockPackedRefs(false)
	if err != nil {
		return err
	}
	if pr == nil {
		return nil
	}
	defer ioutil.CheckClose(pr, &err)

	// Creating the temp file in the same directory as the target file
	// improves our chances for rename operation to be atomic.
	tmp, err := d.fs.TempFile("", tmpPackedRefsPrefix)
	if err != nil {
		return err
	}
	tmpName := tmp.Name()
	defer func() {
		ioutil.CheckClose(tmp, &err)
		_ = d.fs.Remove(tmpName) // don't check err, we might have renamed it
	}()

	s := bufio.NewScanner(pr)
	found := false
	for s.Scan() {
		line := s.Text()
		ref, err := d.processLine(line)
		if err != nil {
			return err
		}

		if ref != nil && ref.Name() == name {
			found = true
			continue
		}

		if _, err := fmt.Fprintln(tmp, line); err != nil {
			return err
		}
	}

	if err := s.Err(); err != nil {
		return err
	}

	if !found {
		return nil
	}

	return d.rewritePackedRefsWhileLocked(tmp, pr)
}

func (d *DotGit) openAndLockPackedRefs(doCreate bool) (pr billy.File, err error) {
	var f billy.File
	defer func() {
		if err != nil && f != nil {
			ioutil.CheckClose(f, &err)
		}
	}()

	// File mode is retrieved from a constant defined in the target specific
	// files (dotgit_rewrite_packed_refs_*). Some modes are not available
	// in all filesystems.
	openFlags := d.openAndLockPackedRefsMode()
	if doCreate {
		openFlags |= os.O_CREATE
	}

	// Keep trying to open and lock the file until we're sure the file
	// didn't change between the open and the lock.
	for {
		f, err = d.fs.OpenFile(packedRefsPath, openFlags, 0600)
		if err != nil {
			if os.IsNotExist(err) && !doCreate {
				return nil, nil
			}

			return nil, err
		}
		fi, err := d.fs.Stat(packedRefsPath)
		if err != nil {
			return nil, err
		}
		mtime := fi.ModTime()

		err = f.Lock()
		if err != nil {
			return nil, err
		}

		fi, err = d.fs.Stat(packedRefsPath)
		if err != nil {
			return nil, err
		}
		if mtime.Equal(fi.ModTime()) {
			break
		}
		// The file has changed since we opened it.  Close and retry.
		err = f.Close()
		if err != nil {
			return nil, err
		}
	}
	return f, nil
}

func (d *DotGit) CountLooseRefs() (int, error) {

}

func (d *DotGit) PackRefs() error {

}
