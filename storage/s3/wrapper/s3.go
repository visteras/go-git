package wrapper

import (
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type s3Client struct {
	sess *session.Session

	options Options
}

func Initialize(endpoint, region string, opts ...Option) (*s3Client, error) {
	options := &Options{}
	for _, o := range opts {
		o(options)
	}

	var cred *credentials.Credentials
	if options.accessKeyId != "" && options.secretAccessKey != "" {
		cred = credentials.NewStaticCredentials(
			options.accessKeyId,
			options.secretAccessKey,
			"",
		)
	}

	sess, err := session.NewSession(
		&aws.Config{
			Region:                        aws.String(region),
			Endpoint:                      aws.String(endpoint),
			Credentials:                   cred,
			CredentialsChainVerboseErrors: aws.Bool(options.verboseCredentials),
		},
	)
	if err != nil {
		return nil, err
	}

	return &s3Client{
		sess:    sess,
		options: *options,
	}, nil
}

func (s *s3Client) PresignRequest(bucket, filename string, duration time.Duration) (string, error) {
	svc := s3.New(s.sess)
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(s.options.prefix + filename),
	})

	return req.Presign(duration)
}

func (s s3Client) Join(elem ...string) string {
	elems := make([]string, len(elem))
	for i, el := range elem {
		elems[i] = strings.Trim(el, "/")
	}
	return "/" + strings.Join(elems, "/")
}
