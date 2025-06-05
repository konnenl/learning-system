package model

import (
	"gorm.io/gorm"
)

type Progress struct {
	gorm.Model
	UserID     uint     `gorm:"uniqueIndex:idx_user_category"`
	CategoryID uint     `gorm:"uniqueIndex:idx_user_category"`
	User       User     `gorm:"foreignKey:UserID;references:ID"`
	Category   Category `gorm:"foreignKey:CategoryID;references:ID"`
	Caterogy   Category `gorm:"foreignKey:CategoryID"`
	Value      float32  `gorm:"default:0"`
}
