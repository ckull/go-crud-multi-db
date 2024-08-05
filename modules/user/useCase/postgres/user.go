package postgres

import (
	models "go-crud/modules/user/model/postgres"
	"go-crud/modules/user/repository/postgres"
	"go-crud/modules/user/useCase/types"
)

type (
	userUsecase struct {
		userRepository postgres.UserRepository
	}
)

func NewUserUsecase(userRepository postgres.UserRepository) types.UserUsecase {
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

func (uc *userUsecase) UpdateUser(id string, update models.User) error {
	return uc.userRepository.UpdateUser(id, update)
}

func (uc *userUsecase) DeleteUser(id string) error {
	return uc.userRepository.DeleteUser(id)
}
