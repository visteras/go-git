package s3

import (
	"github.com/go-git/go-git/v5/plumbing/format/index"
)

type IndexStorage struct {
	s3 s3Git
}

func (i IndexStorage) SetIndex(idx *index.Index) error {
	return i.s3.SetIndex(idx)
}

func (i IndexStorage) Index() (*index.Index, error) {
	return i.s3.Index()
}
