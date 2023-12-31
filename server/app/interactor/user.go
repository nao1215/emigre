// Package interactor is implementation of usecase.
package interactor

import (
	"context"
	"fmt"

	"github.com/nao1215/emigre/server/app/domain/model"
	"github.com/nao1215/emigre/server/app/usecase"
	"github.com/oklog/ulid/v2"
)

var _ usecase.UserCreator = (*UserCreator)(nil)

// UserCreator is an interactor for creating users.
type UserCreator struct{}

// NewUserCreator returns a new UserCreator struct.
func NewUserCreator() *UserCreator {
	return &UserCreator{}
}

// CreateUser creates a user.
func (uc *UserCreator) CreateUser(_ context.Context, input *usecase.UserCreatorInput) (*usecase.UserCreatorOutput, error) {
	user, err := model.NewUser(ulid.Make().String(), input.Name, input.Biography, input.Email)
	if err != nil {
		return nil, fmt.Errorf("can not create user: %w", err)
	}

	// TODO: implement. insert user to database.

	fmt.Println("interactor.UserCreator.CreateUser is not implemented.")
	fmt.Println("user:", user)
	return &usecase.UserCreatorOutput{
		User: user,
	}, nil
}
