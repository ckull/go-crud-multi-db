package postgres

import (
	models "go-crud/modules/user/model/postgres"
	postgresRepo "go-crud/modules/user/repository/postgres"
)

type (
	UserUsecase interface {
		GetUsers() ([]models.User, error)
		CreateUser(user models.User) error
		UpdateUser(id string, update models.User) error
		DeleteUser(id string) error
	}

	userUsecase struct {
		userRepository postgresRepo.UserRepository
	}
)

func NewUserUsecase(userRepository postgresRepo.UserRepository) UserUsecase {
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
