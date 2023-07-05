// Package config load settings from external files or environment variables and manage their values.
package config

import (
	"github.com/caarlos0/env/v9"
)

// S3 is a struct that contains the settings for the S3 storage.
type S3 struct {
	// Bucket is the name of the S3 bucket. The default value is "emigre".
	Bucket string `env:"EMIGRE_S3_BUCKET" envDefault:"emigre"`
	// Region is the name of the AWS region. The default value is "ap-northeast-1".
	Region string `env:"EMIGRE_S3_REGION" envDefault:"ap-northeast-1"`
}

// NewS3 returns a new S3 struct.
func NewS3() (*S3, error) {
	s3 := &S3{}
	if err := env.Parse(s3); err != nil {
		return nil, err
	}
	return s3, nil
}
