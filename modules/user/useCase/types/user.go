package types

import models "go-crud/modules/user/model/mongodb"

type (
	UserUsecase interface {
		GetUsers() ([]models.User, error)
		CreateUser(user models.User) error
		UpdateUser(id string, update models.User) error
		DeleteUser(id string) error
	}
)
