package dotgit

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

type DotGit struct {
	options Options
	path    string
	s3      S3Client
}

// Options holds configuration for the storage.
type Options struct {
}

func New(path string, client S3Client) *DotGit {
	return NewWithOptions(path, client, Options{})
}

// NewWithOptions sets non default configuration options.
// See New for complete help.
func NewWithOptions(path string, client S3Client, o Options) *DotGit {
	return &DotGit{
		options: o,
		s3:      client,
		path:    path,
	}
}
