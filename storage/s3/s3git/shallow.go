package s3git

import (
	"github.com/go-git/go-git/v5/plumbing"
)

func (s *s3Git) SetShallow(commits []plumbing.Hash) error {
	return s.dotgit.SetShallow(commits)
}

func (s *s3Git) Shallow() ([]plumbing.Hash, error) {
	return s.dotgit.Shallow()
}
