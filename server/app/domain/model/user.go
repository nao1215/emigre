package model

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

// User is a user model.
type User struct {
	// ID is user id. ULID（Universally Unique Lexicographically Sortable Identifier）
	ID string `json:"id" validate:"required,min=26,max=26"`
	// Name is user name
	Name string `json:"name" validate:"required,min=1,max=20"`
	// Biography is user self introduction
	Biography string `json:"biography" validate:"required,min=1,max=300"`
	// Email is user email address
	Email string `json:"email" validate:"required,email,max=320"`
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
	var validationErrs validator.ValidationErrors
	if err := validate.Struct(user); errors.As(err, &validationErrs) {
		var errs error
		for _, err := range validationErrs {
			errs = errors.Join(errs, err)
		}
		return nil, errs
	}
	return user, nil
}
