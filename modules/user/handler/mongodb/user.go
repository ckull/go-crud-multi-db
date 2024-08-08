package mongodb

import (
	models "go-crud/modules/user/model/mongodb"
	"go-crud/modules/user/useCase/mongodb"

	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	UserHandler interface {
		GetUsers(c echo.Context) error
		CreateUser(c echo.Context) error
		UpdateUser(c echo.Context) error
		DeleteUser(c echo.Context) error
	}

	userHandler struct {
		userUsecase mongodb.UserUsecase
	}
)

func NewUserHandler(userUsecase mongodb.UserUsecase) UserHandler {
	return &userHandler{
		userUsecase: userUsecase,
	}
}

// @Summary GetUsers
// @Description get users
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User
// @Router /users/ [get]
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
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	updateErr := uh.userUsecase.UpdateUser(objID, user)
	if updateErr != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusCreated)
}

func (uh *userHandler) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	if err := uh.userUsecase.DeleteUser(objID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
