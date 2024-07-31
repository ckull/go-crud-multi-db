package mongodb

import (
	models "go-crud/modules/user/model/mongodb"
	"go-crud/modules/user/repository/mongodb"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	UserUsecase interface {
		GetUsers() ([]models.User, error)
		CreateUser(user models.User) error
		UpdateUser(id primitive.ObjectID, update models.User) error
		DeleteUser(id primitive.ObjectID) error
	}

	userUsecase struct {
		userRepository mongodb.UserRepository
	}
)

func NewUserUsecase(userRepository mongodb.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (uc *userUsecase) GetUsers() ([]models.User, error) {
	return uc.userRepository.GetUsers()
}

func (uc *userUsecase) CreateUser(user models.User) error {
	return uc.userRepository.CreateUser(user)
}

func (uc *userUsecase) UpdateUser(id primitive.ObjectID, update models.User) error {
	return uc.userRepository.UpdateUser(id, update)
}

func (uc *userUsecase) DeleteUser(id primitive.ObjectID) error {
	return uc.userRepository.DeleteUser(id)
}
