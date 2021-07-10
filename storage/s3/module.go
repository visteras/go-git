package s3

import (
	"github.com/go-git/go-git/v5/storage"
)

type ModuleStorage struct {
	s3 s3Git
}

func (m ModuleStorage) Module(name string) (storage.Storer, error) {
	panic("implement me")
}
