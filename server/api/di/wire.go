//go:build wireinject
// +build wireinject

// Package di Inject dependence by wire command.
package di

import (
	"github.com/google/wire"
	"github.com/nao1215/emigre/server/app/interactor"
	"github.com/nao1215/emigre/server/app/usecase"
)

// Emigre is a struct that contains the settings for the Emigre.
type Emigre struct {
	User *User
}

// User is a struct that contains the settings for the user.
type User struct {
	// Creator is an usecase for creating users.
	Creator usecase.UserCreator
}

// NewEmigre returns a new Emigre struct.
func NewEmigre() *Emigre {
	wire.Build(
		interactor.NewUserCreator,
		wire.Bind(new(usecase.UserCreator), new(*interactor.UserCreator)),
		newEmigre,
	)
	return nil
}

// newEmigre returns a new Emigre struct.
func newEmigre(
	uc usecase.UserCreator,
) *Emigre {
	return &Emigre{
		User: &User{
			Creator: uc,
		},
	}
}
