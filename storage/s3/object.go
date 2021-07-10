package s3

import (
	"time"

	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/storer"
)

type ObjectStorage struct {
	s3 s3Git

	//Objects map[plumbing.Hash]plumbing.EncodedObject
	//Commits map[plumbing.Hash]plumbing.EncodedObject
	//Trees   map[plumbing.Hash]plumbing.EncodedObject
	//Blobs   map[plumbing.Hash]plumbing.EncodedObject
	//Tags    map[plumbing.Hash]plumbing.EncodedObject
}

func (o ObjectStorage) DeltaObject(objectType plumbing.ObjectType, hash plumbing.Hash) (plumbing.EncodedObject, error) {
	panic("implement me")
}

func (o ObjectStorage) ObjectPacks() ([]plumbing.Hash, error) {
	panic("implement me")
}

func (o ObjectStorage) DeleteOldObjectPackAndIndex(hash plumbing.Hash, t time.Time) error {
	panic("implement me")
}

func (o ObjectStorage) ForEachObjectHash(f func(plumbing.Hash) error) error {
	panic("implement me")
}

func (o ObjectStorage) LooseObjectTime(hash plumbing.Hash) (time.Time, error) {
	panic("implement me")
}

func (o ObjectStorage) DeleteLooseObject(hash plumbing.Hash) error {
	panic("implement me")
}

func (o ObjectStorage) NewEncodedObject() plumbing.EncodedObject {
	panic("implement me")
}

func (o ObjectStorage) SetEncodedObject(object plumbing.EncodedObject) (plumbing.Hash, error) {
	panic("implement me")
}

func (o ObjectStorage) EncodedObject(objectType plumbing.ObjectType, hash plumbing.Hash) (plumbing.EncodedObject, error) {
	panic("implement me")
}

func (o ObjectStorage) IterEncodedObjects(objectType plumbing.ObjectType) (storer.EncodedObjectIter, error) {
	panic("implement me")
}

func (o ObjectStorage) HasEncodedObject(hash plumbing.Hash) error {
	panic("implement me")
}

func (o ObjectStorage) EncodedObjectSize(hash plumbing.Hash) (int64, error) {
	panic("implement me")
}
