package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nao1215/emigre/version"
)

// HealthResponse is response for GET /health
type HealthResponse struct {
	// Name is server name.
	Name string `json:"name"`
	// Version is repository tag version
	Version string `json:"version"`
	// Revision is git revision
	Revision string `json:"revision"`
}

// newHealthResponse return HealthResponse struct.
func newHealthResponse() *HealthResponse {
	ver := "unknown"
	rev := "unknown"

	if version.TagVersion != "" {
		ver = version.TagVersion
	}
	if version.Revision != "" {
		rev = version.Revision
	}
	return &HealthResponse{
		Name:     "emigre",
		Version:  ver,
		Revision: rev,
	}
}

// health return health response.
func health(c echo.Context) error {
	return c.JSON(http.StatusOK, newHealthResponse())
}
