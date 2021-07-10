package wrapper

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func (s *s3Client) Download(filename string) ([]byte, error) {

	bts := make([]byte, 0)
	buf := aws.NewWriteAtBuffer(bts)
	key := s.Join(s.options.prefix, filename)

	downloader := s3manager.NewDownloader(s.sess)

	_, err := downloader.Download(buf,
		&s3.GetObjectInput{
			Bucket: aws.String(s.options.bucket),
			Key:    aws.String(key),
		})
	if err != nil {
		return nil, ErrDownload.WithError(err).AddDetail("item: %q", key)
	}
	return buf.Bytes(), nil
}
