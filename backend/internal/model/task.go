package model

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Description string `gorm:"unique;not null"`
	Question    string `gorm:"not null"`
	Answer      string `gorm:"not null"`
	CategoryID  uint
	Category    Category `gorm:"foreignKey:CategoryID;references:ID"`
}
