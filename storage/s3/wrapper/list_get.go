package wrapper

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (s *s3Client) GetListObjects(path string) ([]string, error) {
	key := s.Join(s.options.prefix, path)
	params := &s3.ListObjectsInput{
		Bucket: aws.String(s.options.bucket),
		Prefix: aws.String(key),
	}

	svc := s3.New(s.sess)
	resp, err := svc.ListObjects(params)
	if err != nil {
		return nil, ErrGetList.WithError(err).WithBucket(s.options.bucket).AddDetail("path: %q", params.Prefix)
	}
	items := make([]string, 0)
	for _, k := range resp.Contents {
		items = append(items, *k.Key)
	}
	return items, nil
}
