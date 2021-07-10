package s3

import (
	"github.com/go-git/go-git/v5/config"
)

type ConfigStorage struct {
	s3 s3Git
}

func (c ConfigStorage) Config() (*config.Config, error) {
	return c.s3.Config()
}

func (c ConfigStorage) SetConfig(cfg *config.Config) error {
	return c.s3.SetConfig(cfg)
}
