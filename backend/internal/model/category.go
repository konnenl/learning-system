package model

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name  string `gorm:"unique;not null"`
	Level uint   `gorm:"not null"`
	Tasks []Task `gorm:"foreignKey:CategoryID"`
}
