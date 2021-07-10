package s3git

import "github.com/go-git/go-git/v5/plumbing"

const (
	suffix         = ".git"
	packedRefsPath = "packed-refs"
	configPath     = "config"
	indexPath      = "index"
	shallowPath    = "shallow"
	modulePath     = "modules"
	objectsPath    = "objects"
	packPath       = "pack"
	refsPath       = "refs"
	branchesPath   = "branches"
	hooksPath      = "hooks"
	infoPath       = "info"
	remotesPath    = "remotes"
	logsPath       = "logs"
	worktreesPath  = "worktrees"

	tmpPackedRefsPrefix = "._packed-refs"

	packPrefix = "pack-"
	packExt    = ".pack"
	idxExt     = ".idx"
)

type s3Git struct {
	s3 S3Client
}

func (s *s3Git) setRef(fileName string, content string, old *plumbing.Reference) error {
	panic("")
}

func Initialize(s3 S3Client) *s3Git {
	return &s3Git{
		s3: s3,
	}
}
