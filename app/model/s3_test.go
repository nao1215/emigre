// Package model contains the definitions of domain models and business logic.
package model

import (
	"testing"
)

func TestRegionString(t *testing.T) {
	tests := []struct {
		name string
		r    Region
		want string
	}{
		{
			name: "success",
			r:    Region("ap-northeast-1"),
			want: "ap-northeast-1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.String(); got != tt.want {
				t.Errorf("Region.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBucket_Valid(t *testing.T) {
	tests := []struct {
		name string
		b    Bucket
		want bool
	}{
		{
			name: "success",
			b:    Bucket("emigre"),
			want: true,
		},
		{
			name: "failure. bucket name is empty",
			b:    Bucket(""),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Valid(); got != tt.want {
				t.Errorf("Bucket.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBucket_String(t *testing.T) {
	tests := []struct {
		name string
		b    Bucket
		want string
	}{
		{
			name: "success",
			b:    Bucket("emigre"),
			want: "emigre",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.String(); got != tt.want {
				t.Errorf("Bucket.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
