// Package api control emigre api.
package api

import (
	"github.com/labstack/echo/v4"
)

// Run start server.
func Run() error {
	api, err := NewAPI()
	if err != nil {
		return err
	}
	return api.Start(":8080")
}

// API is a structure that aggregates the necessary information for API execution.
type API struct {
	// echo framework. it's manage api handlers.
	*echo.Echo
}

// NewAPI return api struct.
func NewAPI() (*API, error) {
	e := echo.New()
	e.GET("/health", health)

	return &API{
		Echo: e,
	}, nil
}
