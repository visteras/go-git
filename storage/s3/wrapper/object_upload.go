package wrapper

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func (s *s3Client) Upload(filename string, obj []byte) error {

	body := bytes.NewReader(obj)
	key := s.Join(s.options.prefix, filename)

	uploader := s3manager.NewUploader(s.sess)

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s.options.bucket),
		Key:    aws.String(key),
		Body:   body,
	})
	if err != nil {
		return ErrUpload.WithError(err).WithBucket(s.options.bucket).AddDetail("filename: %q", key)
	}

	return nil
}
