// Package repository is an abstraction layer for accessing database.
package repository

import (
	"context"
	"database/sql"
)

// UserCreatorInput is an input struct for UserCreator.
type UserCreatorInput struct {
}

// UserCreatorOutput is an output struct for UserCreator.
type UserCreatorOutput struct {
}

// UserCreator is an interface for creating users.
type UserCreator interface {
	// CreateUser creates a user.
	CreateUser(context.Context, *sql.Tx, *UserCreatorInput) (*UserCreatorOutput, error)
}
