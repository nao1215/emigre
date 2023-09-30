package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nao1215/emigre/server/api/di"
	"github.com/nao1215/emigre/server/app/usecase"
)

// UserController is a controller for /users API.
type UserController struct {
	// Creator is an usecase for creating users.
	Creator usecase.UserCreator
}

// NewUserController returns a new UserController struct.
func NewUserController(u *di.User) *UserController {
	return &UserController{
		Creator: u.Creator,
	}
}

// UserCreatorPayload is a payload for POST /users
type UserCreatorPayload struct {
	// @Description name of user
	Name string `json:"name"`
	// @Description email address of user
	Email string `json:"email"`
	// @Description biography of user
	Biography string `json:"biography"`
}

func (ctrl *UserController) createUser(c echo.Context) error {
	u := new(UserCreatorPayload)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("can not parse payload :%w", err))
	}

	input := &usecase.UserCreatorInput{
		Name:      u.Name,
		Email:     u.Email,
		Biography: u.Biography,
	}
	if _, err := ctrl.Creator.CreateUser(c.Request().Context(), input); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}
