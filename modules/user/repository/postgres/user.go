package postgres

import (
	models "go-crud/modules/user/model/postgres"

	"gorm.io/gorm"
)

type (
	UserRepository interface {
		GetUsers() ([]models.User, error)
		CreateUser(user models.User) error
		GetUserByID(id string) (models.User, error)
		UpdateUser(id string, update models.User) error
		DeleteUser(id string) error
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	result := r.db.Find(&users)
	return users, result.Error
}

func (r *userRepository) CreateUser(user models.User) error {
	err := r.db.Create(&user).Error
	return err
}

func (r *userRepository) GetUserByID(id string) (models.User, error) {
	var user models.User
	err := r.db.First(&user.ID, id).Error
	return user, err
}

func (r *userRepository) UpdateUser(id string, update models.User) error {
	var user models.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return err
	}

	// Apply the updates
	user.Name = update.Name
	user.Email = update.Email

	return r.db.Save(&user).Error
}

func (r *userRepository) DeleteUser(id string) error {
	var user models.User
	return r.db.Delete(&user, id).Error
}
