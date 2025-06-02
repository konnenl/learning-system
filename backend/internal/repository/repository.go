package repository

import (
	"github.com/konnenl/learning-system/internal/model"
	"gorm.io/gorm"
)

type Repository struct {
	User    UserRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User:    newUserRepository(db),
	}
}

type UserRepository interface {
	Create(user *model.User) (uint, error)
	Authenticate(email string, password string) (*model.User, error)
	GetByID(id uint) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
}