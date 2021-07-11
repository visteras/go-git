package s3git

import (
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/storer"
)

func (s *s3Git) SetReference(r *plumbing.Reference) error {
	return s.dotgit.SetRef(r, nil)
}
func (s *s3Git) CheckAndSetReference(r, old *plumbing.Reference) error {
	return s.dotgit.SetRef(r, old)
}
func (s *s3Git) Reference(name plumbing.ReferenceName) (*plumbing.Reference, error) {
	return s.dotgit.Ref(name)
}
func (s *s3Git) IterReferences() (storer.ReferenceIter, error) {
	refs, err := s.dotgit.Refs()
	if err != nil {
		return nil, err
	}

	return storer.NewReferenceSliceIter(refs), nil
}
func (s *s3Git) RemoveReference(n plumbing.ReferenceName) error {
	return s.dotgit.RemoveRef(n)
}
func (s *s3Git) CountLooseRefs() (int, error) {
	return s.dotgit.CountLooseRefs()
}
func (s *s3Git) PackRefs() error {
	return s.dotgit.PackRefs()
}
