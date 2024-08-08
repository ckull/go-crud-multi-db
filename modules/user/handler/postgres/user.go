package postgres

import (
	models "go-crud/modules/user/model/postgres"
	useCase "go-crud/modules/user/useCase/postgres"

	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	UserHandler interface {
		GetUsers(c echo.Context) error
		CreateUser(c echo.Context) error
		UpdateUser(c echo.Context) error
		DeleteUser(c echo.Context) error
	}

	userHandler struct {
		userUsecase useCase.UserUsecase
	}
)

func NewUserHandler(userUsecase useCase.UserUsecase) UserHandler {
	return &userHandler{
		userUsecase: userUsecase,
	}
}

func (uh *userHandler) GetUsers(c echo.Context) error {
	users, err := uh.userUsecase.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, users)
}

func (uh *userHandler) CreateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	err := uh.userUsecase.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusCreated)
}

func (uh *userHandler) UpdateUser(c echo.Context) error {
	id := c.Param("id")

	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	updateErr := uh.userUsecase.UpdateUser(id, user)
	if updateErr != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": updateErr.Error()})
	}

	return c.NoContent(http.StatusCreated)
}

func (uh *userHandler) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	if err := uh.userUsecase.DeleteUser(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
