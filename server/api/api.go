// Package api control emigre api.
package api

import (
	"github.com/labstack/echo/v4"
	// Import docs for Swagger documentation generation
	_ "github.com/nao1215/emigre/docs"
	"github.com/nao1215/emigre/server/api/di"
	echoSwagger "github.com/swaggo/echo-swagger"
)

//	@title			Emigre API Swagger
//	@version		0.0.1
//	@description	This is a emigre API server swagger.
//	@termsOfService	https://github.com/nao1215/emigre

//	@contact.name	Naohiro CHIKAMATSU (Author)
//	@contact.url	https://github.com/nao1215/emigre/server/issues
//	@contact.email	n.chika156@gmail

//	@license.name	MIT License
//	@license.url	https://github.com/nao1215/emigre/server/blob/main/LICENSE

//	@host		https://nao1215.github.io/emigre/html/index.html
//	@BasePath	/v1

// Run start server.
func Run() error {
	return NewAPI().Start(":8080")
}

// API is a structure that aggregates the necessary information for API execution.
type API struct {
	// echo framework. it's manage api handlers.
	*echo.Echo
	// emigre is a struct that contains the settings for the Emigre.
	emigre *di.Emigre
	// healthController is a controller for /health API.
	healthController *HealthController
	// userController is a controller for /users API.
	userController *UserController
}

// NewAPI return api struct.
func NewAPI() *API {
	api := new(API)
	api.Echo = echo.New()
	api.emigre = di.NewEmigre()
	api.setControllers()
	api.route()

	return api
}

// setControllers set controllers.
// This method is called before routing.
func (a *API) setControllers() {
	a.healthController = NewHealthController()
	a.userController = NewUserController(a.emigre.User)
}

// route set routing.
// This method is called after setControllers.
func (a *API) route() {
	a.GET("/v1/health", a.healthController.health)
	a.POST("/v1/users", a.userController.createUser)
	a.GET("/swagger/*", echoSwagger.WrapHandler)
}
