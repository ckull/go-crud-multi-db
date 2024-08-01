package postgres

import "gorm.io/gorm"

type (
	UserRepository interface {
	}

	userRepository struct {
		db *gorm.DB
	}
)
