package s3git

import (
	"bufio"
	"bytes"
	"fmt"

	"github.com/go-git/go-git/v5/plumbing"
)

func (s *s3Git) SetShallow(commits []plumbing.Hash) error {
	buf := bytes.NewBuffer(nil)

	for _, h := range commits {
		if _, err := fmt.Fprintf(buf, "%s\n", h); err != nil {
			return err
		}
	}

	return s.s3.Upload(shallowPath, buf.Bytes())
}

func (s *s3Git) Shallow() ([]plumbing.Hash, error) {
	b, err := s.s3.Download(shallowPath)
	if err != nil {
		return nil, err
	}

	var hash []plumbing.Hash
	f := bytes.NewReader(b)

	scn := bufio.NewScanner(f)
	for scn.Scan() {
		hash = append(hash, plumbing.NewHash(scn.Text()))
	}

	return hash, scn.Err()
}
