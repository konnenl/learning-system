package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/konnenl/learning-system/internal/model"
)

func New() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("./data/data.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(1)
	return db, nil
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&model.User{},
		&model.Category{},
		&model.Task{},
		&model.Progress{},
		&model.Word{},
	)
	return err
}
