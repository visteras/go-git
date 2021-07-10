package s3git

type S3Client interface {
	CreateBucket(name string) error

	Upload(filename string, body []byte) error
	Download(filename string) ([]byte, error)
	Delete(path string) error

	Join(elem ...string) string
}
