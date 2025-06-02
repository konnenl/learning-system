package repository

import (
	"github.com/konnenl/learning-system/internal/model"
	"gorm.io/gorm"
)

type Repository struct {
	User    UserRepository
	Word WordRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User:    newUserRepository(db),
		Word:    newWordRepository(db),
	}
}

type UserRepository interface {
	Create(user *model.User) (uint, error)
	Authenticate(email string, password string) (*model.User, error)
	GetByID(id uint) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	GetLevel(id uint) (string, error)
}

type WordRepository interface{
	GetWords() ([]*model.Word, error)
}