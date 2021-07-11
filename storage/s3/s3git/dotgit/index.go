package dotgit

import (
	"bufio"
	"bytes"

	"github.com/go-git/go-git/v5/plumbing/format/index"
)

func (d *DotGit) SetIndex(idx *index.Index) error {
	bw := bytes.NewBuffer(nil)

	e := index.NewEncoder(bw)
	err := e.Encode(idx)
	if err != nil {
		return err
	}

	return d.s3.Upload(indexPath, bw.Bytes())
}
func (d *DotGit) Index() (*index.Index, error) {
	idx := &index.Index{
		Version: 2,
	}

	b, err := d.s3.Download(indexPath)
	if err != nil {
		return nil, err
	}

	f := bytes.NewBuffer(b)
	decoder := index.NewDecoder(bufio.NewReader(f))
	err = decoder.Decode(idx)
	return idx, err
}
