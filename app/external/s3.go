// Package external is the implementation for accessing external services.
package external

import (
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/nao1215/emigre/app/service"
	"github.com/nao1215/emigre/config"
)

// S3 is an implementation for FileDownloader.
type S3 struct {
	*s3manager.Downloader
}

var _ service.FileDownloder = &S3{}

// NewS3 returns a new s3 struct.
func NewS3(config config.S3) *S3 {
	awsConfig := &aws.Config{
		Region: aws.String(config.Region.String()),
	}

	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable, // Ref. ~/.aws/config
		Config:            *awsConfig,
	}))

	downloader := s3manager.NewDownloader(session, func(d *s3manager.Downloader) {
		d.BufferProvider = s3manager.NewPooledBufferedWriterReadFromProvider(5 * 1024 * 1024) // 5MB
	})
	return &S3{downloader}
}

// DownloadFile downloads a file from external storage.
func (s *S3) DownloadFile(_ context.Context, input *service.FileDownloderInput) (*service.FileDownloderOutput, error) {
	buf := aws.NewWriteAtBuffer([]byte{})
	objInput := &s3.GetObjectInput{
		Bucket: aws.String(input.Config.Bucket.String()),
		Key:    aws.String(input.Key),
	}

	if _, err := s.Download(buf, objInput); err != nil {
		return nil, err
	}
	return &service.FileDownloderOutput{
		Buffer: bytes.NewBuffer(buf.Bytes()),
	}, nil
}
