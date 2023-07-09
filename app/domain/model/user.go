package model

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

// User is a user model.
type User struct {
	// ID is user id. ULID（Universally Unique Lexicographically Sortable Identifier）
	ID string
	// Name is user name
	Name string `validate:"required,min=1,max=20"`
	// Biography is user self introduction
	Biography string `validate:"required,min=1,max=300"`
	// Email is user email address
	Email string `validate:"required,email"`
}

// NewUser is a constructor of User.
func NewUser(id, name, biography, email string) (*User, error) {
	user := &User{
		ID:        id,
		Name:      name,
		Biography: biography,
		Email:     email,
	}

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		validationErrs := err.(validator.ValidationErrors)
		var errs error
		for _, err := range validationErrs {
			errs = errors.Join(errs, err)
		}
		return nil, errs
	}
	return user, nil
}
