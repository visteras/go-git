package s3

import (
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/storer"
)

type ReferenceStorage struct {
	s3 s3Git
}

func (r ReferenceStorage) SetReference(reference *plumbing.Reference) error {
	return r.s3.SetReference(reference)
}

func (r ReferenceStorage) CheckAndSetReference(new, old *plumbing.Reference) error {
	return r.s3.CheckAndSetReference(new, old)
}

func (r ReferenceStorage) Reference(name plumbing.ReferenceName) (*plumbing.Reference, error) {
	return r.s3.Reference(name)
}

func (r ReferenceStorage) IterReferences() (storer.ReferenceIter, error) {
	return r.s3.IterReferences()
}

func (r ReferenceStorage) RemoveReference(name plumbing.ReferenceName) error {
	return r.s3.RemoveReference(name)
}

func (r ReferenceStorage) CountLooseRefs() (int, error) {
	return r.s3.CountLooseRefs()
}

func (r ReferenceStorage) PackRefs() error {
	return r.s3.PackRefs()
}
