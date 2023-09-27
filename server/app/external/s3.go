// Package external is the implementation for accessing external services.
package external

import (
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/nao1215/emigre/server/app/domain/service"
	"github.com/nao1215/emigre/server/config"
)

// S3Downloader is an implementation for FileDownloader.
type S3Downloader struct {
	*s3manager.Downloader
}

var _ service.FileDownloder = &S3Downloader{}

// downloadBufferSize is the buffer size for downloading files. It's 5MB.
const downloadBufferSize = 5 * 1024 * 1024

// NewS3Downloader returns a new S3Downloader struct.
func NewS3Downloader(config config.S3) *S3Downloader {
	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable, // Ref. ~/.aws/config
		Config:            aws.Config{Region: aws.String(config.Region.String())},
	}))
	downloader := s3manager.NewDownloader(session, func(d *s3manager.Downloader) {
		d.BufferProvider = s3manager.NewPooledBufferedWriterReadFromProvider(downloadBufferSize)
	})
	return &S3Downloader{downloader}
}

// DownloadFile downloads a file from S3.
func (s *S3Downloader) DownloadFile(_ context.Context, input *service.FileDownloderInput) (*service.FileDownloderOutput, error) {
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

// S3Uploader is an implementation for FileUploader.
type S3Uploader struct {
	*s3manager.Uploader
}

var _ service.FileUploader = &S3Uploader{}

// NewS3Uploader returns a new S3Uploader struct.
func NewS3Uploader(config config.S3) *S3Uploader {
	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable, // Ref. ~/.aws/config
		Config:            aws.Config{Region: aws.String(config.Region.String())},
	}))
	return &S3Uploader{s3manager.NewUploader(session)}
}

// UploadFile uploads a file to S3.
func (s *S3Uploader) UploadFile(_ context.Context, input *service.FileUploaderInput) (*service.FileUploaderOutput, error) {
	uploadInput := &s3manager.UploadInput{
		Bucket: aws.String(input.Config.Bucket.String()),
		Body:   aws.ReadSeekCloser(input.Data),
		Key:    aws.String(input.Key),
		// TODO: Set ContentType
	}

	if _, err := s.Upload(uploadInput); err != nil {
		return nil, err
	}
	return &service.FileUploaderOutput{}, nil
}
