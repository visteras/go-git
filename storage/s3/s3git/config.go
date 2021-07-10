package s3git

import (
	"bytes"

	"github.com/go-git/go-git/v5/config"
)

func (s *s3Git) SetConfig(cfg *config.Config) error {
	if err := cfg.Validate(); err != nil {
		return err
	}

	b, err := cfg.Marshal()
	if err != nil {
		return err
	}

	return s.s3.Upload(configPath, b)
}

func (s *s3Git) Config() (*config.Config, error) {
	b, err := s.s3.Download(configPath)
	if err != nil {
		return nil, err
	}

	f := bytes.NewBuffer(b)

	return config.ReadConfig(f)
}
