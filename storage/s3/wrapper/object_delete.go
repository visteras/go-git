package wrapper

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (s *s3Client) Delete(filename string) error {
	svc := s3.New(s.sess)

	key := s.Join(s.options.prefix, filename)

	_, err := svc.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(s.options.bucket), Key: aws.String(key)})
	if err != nil {
		return ErrDeleted.WithError(err).WithBucket(s.options.bucket).AddDetail("item: %q", key)
	}

	err = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(s.options.bucket),
		Key:    aws.String(key),
	})
	return ErrDeleted.WithError(err).WithMsg("couldn't wait until the object not exists").WithBucket(s.options.bucket).AddDetail("item: %q", key)
}
