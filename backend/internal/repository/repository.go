package repository

import (
	"github.com/konnenl/learning-system/internal/model"
	"gorm.io/gorm"
)

type Repository struct {
	User     UserRepository
	Word     WordRepository
	Category CategoryRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User:     newUserRepository(db),
		Word:     newWordRepository(db),
		Category: newCategoryRepository(db),
	}
}

type UserRepository interface {
	Create(user *model.User) (uint, error)
	Authenticate(email string, password string) (*model.User, error)
	GetByID(id uint) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	GetLevel(id uint) (string, error)
	UpdateLevel(id uint, level string) error
	UpgrateLevel(id uint) (string, error)
	GetAllAdmin() ([]model.User, error)
}

type WordRepository interface {
	GetWords() ([]*model.Word, error)
	GetLevels(words_id []uint) ([]string, error)
}

type CategoryRepository interface {
	GetAllCategories() ([]model.Category, error)
	Create(category *model.Category) (uint, error)
	Delete(categoryID uint) error
	CreateTask(task *model.Task) (uint, error)
	DeleteTask(taskID uint) error
	GetCategoryTasks(id uint) (model.Category, error)
	GetNextCategory(id uint) (model.Category, error)
	UpdateProgress(userID uint, categoryID uint, value float32) error
}
