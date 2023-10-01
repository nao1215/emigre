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
	//	@Description	name of user
	Name string `json:"name"`
	//	@Description	email address of user
	Email string `json:"email"`
	//	@Description	biography of user
	Biography string `json:"biography"`
}

// createUser create user.
//
//	@Summary		create user.
//	@Description	This API crate user record. If the user already exists, it returns an error.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			name		body		string	true	"user name. 0 < len < 21"
//	@Param			biography	body		string	true	"user biography. max length is 300"
//	@Param			email		body		string	true	"email address. max length is 320"
//	@Success		201			{object}	model.User
//	@Failure		500			{object}	string
//	@Router			/users [post]
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
	output, err := ctrl.Creator.CreateUser(c.Request().Context(), input)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, output.User)
}
