// Package usecase is an abstraction layer for business logic.
package usecase

import (
	"context"

	"github.com/nao1215/emigre/server/app/domain/model"
)

// UserCreatorInput is an input struct for UserCreator.
type UserCreatorInput struct {
	// Name is the user's name.
	Name string
	// Email is the user's email address.
	Email string
	// Biography is the user's self introduction.
	Biography string
}

// UserCreatorOutput is an output struct for UserCreator.
type UserCreatorOutput struct {
	User *model.User
}

// UserCreator is an interface for creating users.
type UserCreator interface {
	// CreateUser creates a user.
	CreateUser(context.Context, *UserCreatorInput) (*UserCreatorOutput, error)
}
