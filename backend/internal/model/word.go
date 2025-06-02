package model

import (
	"gorm.io/gorm"
)

type Word struct {
	gorm.Model
	Word  string `gorm:"unique;not null"`
	Level string   `gorm:"not null"`
}
