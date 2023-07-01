package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nao1215/emigre/version"
)

// HealthResponse is response for GET /health
type HealthResponse struct {
	// @Description Name is server name.
	Name string `json:"name"`
	// @Description Version is repository tag version
	Version string `json:"version"`
	// @Description Revision is git revision
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
// @Summary Get server health information
// @Description This API is for checking server health. The api return server name, version and revision.
// @Success 200 {object} HealthResponse
// @Router /health [get]
func health(c echo.Context) error {
	return c.JSON(http.StatusOK, newHealthResponse())
}
