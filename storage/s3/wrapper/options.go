package wrapper

type Options struct {
	prefix             string
	accessKeyId        string
	secretAccessKey    string
	verboseCredentials bool
	bucket             string
}

type Option func(opts *Options)

func SetPrefix(prefix string) Option {
	return func(o *Options) {
		o.prefix = prefix
	}
}

func SetBucket(bucket string) Option {
	return func(o *Options) {
		o.bucket = bucket
	}
}

func SetAccessKeyId(accessKeyId string) Option {
	return func(o *Options) {
		o.accessKeyId = accessKeyId
	}
}

func SetSecretAccessKey(secretAccessKey string) Option {
	return func(o *Options) {
		o.secretAccessKey = secretAccessKey
	}
}

func SetVerboseCredentials(verboseCredentials bool) Option {
	return func(o *Options) {
		o.verboseCredentials = verboseCredentials
	}
}
