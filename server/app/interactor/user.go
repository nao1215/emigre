// Package interactor is implementation of usecase.
package interactor

import (
	"context"
	"fmt"

	"github.com/nao1215/emigre/server/app/usecase"
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
	fmt.Println("interactor.UserCreator.CreateUser is not implemented.")
	fmt.Println("input:", input)
	return nil, nil
}
