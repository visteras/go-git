package s3git

import (
	"bufio"
	"bytes"

	"github.com/go-git/go-git/v5/plumbing/format/index"
)

func (s *s3Git) SetIndex(idx *index.Index) error {
	bw := bytes.NewBuffer(nil)

	e := index.NewEncoder(bw)
	err := e.Encode(idx)
	if err != nil {
		return err
	}

	return s.s3.Upload(indexPath, bw.Bytes())
}
func (s *s3Git) Index() (*index.Index, error) {
	idx := &index.Index{
		Version: 2,
	}

	b, err := s.s3.Download(indexPath)
	if err != nil {
		return nil, err
	}

	f := bytes.NewBuffer(b)
	d := index.NewDecoder(bufio.NewReader(f))
	err = d.Decode(idx)
	return idx, err
}
