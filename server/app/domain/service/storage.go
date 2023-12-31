// Package service is an abstraction layer for accessing external services.
package service

import (
	"bytes"
	"context"
	"io"

	"github.com/nao1215/emigre/server/config"
)

// FileDownloderInput is an input struct for FileDownloader.
type FileDownloderInput struct {
	// Config is S3 Config.
	Config config.S3
	// Key is the S3 key
	Key string
}

// FileDownloderOutput is an output struct for FileDownloader.
type FileDownloderOutput struct {
	// Buffer is the downloaded data.
	*bytes.Buffer
}

// FileDownloder is an interface for download file to external storage.
type FileDownloder interface {
	// DownloadFile downloads a file from external storage.
	DownloadFile(context.Context, *FileDownloderInput) (*FileDownloderOutput, error)
}

// FileUploaderInput is an input struct for FileUploader.
type FileUploaderInput struct {
	// Config is S3 Config.
	Config config.S3
	// Key is the S3 key
	Key string
	// Data is the data to upload.
	Data io.Reader
}

// FileUploaderOutput is an output struct for FileUploader.
type FileUploaderOutput struct {
}

// FileUploader is an interface for uploading files to external storage.
type FileUploader interface {
	// UploadFile uploads a file from external storage.
	UploadFile(context.Context, *FileUploaderInput) (*FileUploaderOutput, error)
}
