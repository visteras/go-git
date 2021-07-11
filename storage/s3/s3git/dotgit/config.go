package dotgit

import (
	"bytes"

	"github.com/go-git/go-git/v5/config"
)

func (d *DotGit) SetConfig(cfg *config.Config) error {
	if err := cfg.Validate(); err != nil {
		return err
	}

	b, err := cfg.Marshal()
	if err != nil {
		return err
	}

	return d.s3.Upload(configPath, b)
}

func (d *DotGit) Config() (*config.Config, error) {
	b, err := d.s3.Download(configPath)
	if err != nil {
		return nil, err
	}

	f := bytes.NewBuffer(b)

	return config.ReadConfig(f)
}
