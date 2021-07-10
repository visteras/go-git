package s3git

import (
	"fmt"

	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/storer"
)

func (s *s3Git) SetReference(r *plumbing.Reference) error {
	var content string
	switch r.Type() {
	case plumbing.SymbolicReference:
		content = fmt.Sprintf("ref: %s\n", r.Target())
	case plumbing.HashReference:
		content = fmt.Sprintln(r.Hash().String())
	}

	fileName := r.Name().String()

	return s.setRef(fileName, content, nil)
}
func (s *s3Git) CheckAndSetReference(new, old *plumbing.Reference) error {
	panic("")
}
func (s *s3Git) Reference(name plumbing.ReferenceName) (*plumbing.Reference, error) {
	panic("")
}
func (s *s3Git) IterReferences() (storer.ReferenceIter, error) {
	panic("")
}
func (s *s3Git) RemoveReference(name plumbing.ReferenceName) error {
	panic("")
}
func (s *s3Git) CountLooseRefs() (int, error) {
	panic("")
}
func (s *s3Git) PackRefs() error {
	panic("")
}
