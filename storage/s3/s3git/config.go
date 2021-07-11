package s3git

import (
	"github.com/go-git/go-git/v5/config"
)

func (s *s3Git) SetConfig(cfg *config.Config) error {
	return s.dotgit.SetConfig(cfg)
}

func (s *s3Git) Config() (*config.Config, error) {
	return s.dotgit.Config()
}
