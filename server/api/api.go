// Package api control emigre api.
package api

import (
	"github.com/labstack/echo/v4"
	// Import docs for Swagger documentation generation
	_ "github.com/nao1215/emigre/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Emigre API Swagger
// @version 0.0.1
// @description This is a emigre API server swagger.
// @termsOfService https://github.com/nao1215/emigre

// @contact.name Naohiro CHIKAMATSU (Author)
// @contact.url https://github.com/nao1215/emigre/server/issues
// @contact.email n.chika156@gmail

// @license.name MIT License
// @license.url https://github.com/nao1215/emigre/server/blob/main/LICENSE

// @host https://nao1215.github.io/emigre/html/index.html
// @BasePath /v1

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

	// TODO: only add in debug mode
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	return &API{
		Echo: e,
	}, nil
}
