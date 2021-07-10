package wrapper

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (s *s3Client) GetList(bucket, path string) ([]string, error) {

	params := &s3.ListObjectsInput{
		Bucket: aws.String(bucket),
		Prefix: aws.String(s.options.prefix + path),
	}

	svc := s3.New(s.sess)
	resp, err := svc.ListObjects(params)
	if err != nil {
		return nil, ErrGetList.WithError(err).WithBucket(bucket).AddDetail("path: %q", params.Prefix)
	}
	items := make([]string, 0)
	for _, key := range resp.Contents {
		items = append(items, *key.Key)
	}
	return items, nil

}
