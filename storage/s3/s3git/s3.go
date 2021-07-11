package s3git

import "github.com/go-git/go-git/v5/storage/s3/s3git/dotgit"

type s3Git struct {
	s3     dotgit.S3Client
	dotgit *dotgit.DotGit
}

func Initialize(s3 dotgit.S3Client) *s3Git {
	dg := dotgit.New("", s3)
	return &s3Git{
		s3:     s3,
		dotgit: dg,
	}
}
