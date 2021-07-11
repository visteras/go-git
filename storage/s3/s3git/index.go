package s3git

import (
	"github.com/go-git/go-git/v5/plumbing/format/index"
)

func (s *s3Git) SetIndex(idx *index.Index) error {
	return s.dotgit.SetIndex(idx)
}
func (s *s3Git) Index() (*index.Index, error) {
	return s.dotgit.Index()
}
