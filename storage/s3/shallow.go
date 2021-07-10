package s3

import (
	"github.com/go-git/go-git/v5/plumbing"
)

type ShallowStorage struct {
	s3 s3Git
}

func (s ShallowStorage) SetShallow(commits []plumbing.Hash) error {
	currents, err := s.Shallow()
	if err != nil {
		return err
	}
	currents = append(currents, commits...)
	return s.s3.SetShallow(currents)
}

func (s ShallowStorage) Shallow() ([]plumbing.Hash, error) {
	return s.s3.Shallow()
}
