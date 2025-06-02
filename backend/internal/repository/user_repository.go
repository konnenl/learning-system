package repository

import (
	"errors"
	"github.com/konnenl/learning-system/internal/model"
	"gorm.io/gorm"
	"strings"
)

type userRepository struct {
	db *gorm.DB
}

func newUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *model.User) (uint, error) {
	if err := r.db.Create(&user).Error; err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed: users.email") {
			return 0, errors.New("email already exist")
		}
		return 0, err
	}
	return user.ID, nil
}

func (r *userRepository) Authenticate(email string, password string) (*model.User, error) {
	user, err := r.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	if model.CheckPassword(user.Password, password) != nil {
		return nil, errors.New("invalid password")
	}
	return user, nil
}

func (r *userRepository) GetByID(id uint) (*model.User, error) {
	var user model.User
	if err := r.db.Where("user_id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetLevel(id uint) (string, error) {
	var level string
	err := r.db.
		Model(&model.User{}).
		Where("id = ?", id).
		Select("level").
		Scan(&level).
		Error
	if err != nil {
		return "", err
	}
	return level, nil
}

func (r *userRepository) UpdateLevel(id uint, level string) error {
	err := r.db.
		Model(&model.User{}).
		Where("id = ?", id).
		Update("level", level).
		Error
	return err
}
