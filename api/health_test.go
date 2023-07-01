package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/nao1215/emigre/version"
	"github.com/stretchr/testify/assert"
)

func TestHealthHandler(t *testing.T) {
	t.Parallel()

	t.Run("return health response", func(t *testing.T) {
		t.Parallel()

		e := echo.New()
		e.GET("/health", health)

		req := httptest.NewRequest(http.MethodGet, "/health", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := health(c) // Kick handler

		// Check response
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		// Check json response
		want := newHealthResponse()
		expected, err := json.Marshal(want)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, fmt.Sprintf("%s\n", expected), rec.Body.String())
	})
}

// This test is not parallel because it changes the global variable.
func TestNewHealthResponse(t *testing.T) {
	type args struct {
		tagVersion string
		revision   string
	}
	tests := []struct {
		name string
		args args
		want *HealthResponse
	}{
		{
			name: "unknown version",
			args: args{
				tagVersion: "",
				revision:   "rev",
			},
			want: &HealthResponse{
				Name:     "emigre",
				Version:  "unknown",
				Revision: "rev",
			},
		},
		{
			name: "unknown revision",
			args: args{
				tagVersion: "v0.0.1",
				revision:   "",
			},
			want: &HealthResponse{
				Name:     "emigre",
				Version:  "v0.0.1",
				Revision: "unknown",
			},
		},
		{
			name: "unknown version and revision",
			args: args{
				tagVersion: "",
				revision:   "",
			},
			want: &HealthResponse{
				Name:     "emigre",
				Version:  "unknown",
				Revision: "unknown",
			},
		},
		{
			name: "Get version and revision",
			args: args{
				tagVersion: "v0.0.0",
				revision:   "rev",
			},
			want: &HealthResponse{
				Name:     "emigre",
				Version:  "v0.0.0",
				Revision: "rev",
			},
		},
	}
	for _, tt := range tests {
		version.TagVersion = tt.args.tagVersion
		version.Revision = tt.args.revision

		t.Run(tt.name, func(t *testing.T) {
			if got := newHealthResponse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newHealthResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
