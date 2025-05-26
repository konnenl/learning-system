package model

import (
	"gorm.io/gorm"
)

type Progress struct {
	gorm.Model
	UserID     uint
	CategoryID uint
	User       User     `gorm:"foreignKey:UserID;references:ID"` 
    Category   Category `gorm:"foreignKey:CategoryID;references:ID"`
	Caterogy   Category `gorm:"foreignKey:CategoryID"`
	Progress   float64 `gorm:"default:0"`
}
