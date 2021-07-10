package s3

import (
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/format/index"
	"github.com/go-git/go-git/v5/plumbing/storer"
	"github.com/go-git/go-git/v5/storage/s3/s3git"
	"github.com/go-git/go-git/v5/storage/s3/wrapper"
)

type s3Git interface {
	SetConfig(*config.Config) error
	Config() (*config.Config, error)

	SetIndex(*index.Index) error
	Index() (*index.Index, error)

	SetShallow([]plumbing.Hash) error
	Shallow() ([]plumbing.Hash, error)

	SetReference(*plumbing.Reference) error
	CheckAndSetReference(*plumbing.Reference, *plumbing.Reference) error
	Reference(plumbing.ReferenceName) (*plumbing.Reference, error)
	IterReferences() (storer.ReferenceIter, error)
	RemoveReference(plumbing.ReferenceName) error
	CountLooseRefs() (int, error)
	PackRefs() error
}

type Storage struct {
	s3 s3Git

	ConfigStorage
	ShallowStorage
	IndexStorage

	ObjectStorage
	ReferenceStorage

	ModuleStorage
}

// NewStorage returns a new Storage backed by a given `fs.Filesystem` and cache.
func NewStorage(endpoint, region string) (*Storage, error) {
	return NewStorageWithOptions(endpoint, region)
}

// NewStorageWithOptions returns a new Storage with extra options,
// backed by a given `fs.Filesystem` and cache.
func NewStorageWithOptions(endpoint, region string, s3opts ...wrapper.Option) (*Storage, error) {
	s3, err := wrapper.Initialize(endpoint, region, s3opts...)
	if err != nil {
		return nil, err
	}
	s3Git := s3git.Initialize(s3)
	return &Storage{
		s3: s3Git,

		ObjectStorage:    ObjectStorage{s3: s3Git},
		ReferenceStorage: ReferenceStorage{s3: s3Git},
		IndexStorage:     IndexStorage{s3: s3Git},
		ShallowStorage:   ShallowStorage{s3: s3Git},
		ConfigStorage:    ConfigStorage{s3: s3Git},
		ModuleStorage:    ModuleStorage{s3: s3Git},
	}, nil
}

// Init initializes .git directory
func (s *Storage) Init() error {

	return nil
}
