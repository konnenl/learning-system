package model

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name  string `gorm:"unique;not null"`
	Level string `gorm:"not null"`
	Tasks []Task `gorm:"foreignKey:CategoryID;constraint:OnDelete:CASCADE"`
}
