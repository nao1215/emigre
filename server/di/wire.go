//go:build wireinject
// +build wireinject

// Package di Inject dependence by wire command.
package di

import (
	"github.com/google/wire"
	"github.com/nao1215/emigre/server/api"
	"github.com/nao1215/emigre/server/app/interactor"
	"github.com/nao1215/emigre/server/app/usecase"
)

// Emigre is a struct that contains the settings for the Emigre.
type Emigre struct {
	// UserController is a controller for users apis.
	UserController api.UserController
	// HealthController is a controller for health apis.
	HealthController api.HealthController
}

// NewEmigre returns a new Emigre struct.
func NewEmigre() (*Emigre, error) {
	wire.Build(
		interactor.NewUserCreator,
		wire.Bind(new(usecase.UserCreator), new(*interactor.UserCreator)),
		api.NewUserController,
		api.NewHealthController,
		newEmigre,
	)
	return nil, nil
}

// newEmigre returns a new Emigre struct.
func newEmigre(
	uc api.UserController,
	hc api.HealthController,
) *Emigre {
	return &Emigre{
		UserController:   uc,
		HealthController: hc,
	}
}
