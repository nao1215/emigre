package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nao1215/emigre/server/version"
)

// HealthController is a controller for /health API.
type HealthController struct{}

// NewHealthController returns a new HealthController struct.
func NewHealthController() *HealthController {
	return &HealthController{}
}

// HealthResponse is response for GET /health
type HealthResponse struct {
	//	@Description	Name is server name.
	Name string `json:"name"`
	//	@Description	Version is repository tag version
	Version string `json:"version"`
	//	@Description	Revision is git revision
	Revision string `json:"revision"`
}

// health return health response.
//
//	@Summary		Get server health information
//	@Description	This API is for checking server health. The api return server name, version and revision.
//	@Tags			health
//	@Produce		json
//	@Success		200	{object}	HealthResponse
//	@Router			/health [get]
func (ctrl *HealthController) health(c echo.Context) error {
	return c.JSON(http.StatusOK, ctrl.newHealthResponse())
}

// newHealthResponse return HealthResponse struct.
func (ctrl *HealthController) newHealthResponse() *HealthResponse {
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
