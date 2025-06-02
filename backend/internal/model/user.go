package model

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string     `gorm:"unique;not null"`
	Password string     `gorm:"not null"`
	Fullname string     `gorm:"not null"`
	Role     string     `gorm:"not null"`
	Level    *string    `gorm:"default:''"`
	Progress []Progress `gorm:"foreignKey:UserID"`
}

func HashPassword(password string) (string, error) {
	if len(password) == 0 {
		return "", errors.New("password should not be empty")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func CheckPassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
