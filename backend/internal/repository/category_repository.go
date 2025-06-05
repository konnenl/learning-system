package repository

import (
	"github.com/konnenl/learning-system/internal/model"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func newCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) GetAllCategories() ([]model.Category, error) {
	var categories []model.Category
	err := r.db.
		Find(&categories).
		Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepository) Create(category *model.Category) (uint, error) {
	if err := r.db.Create(&category).Error; err != nil {
		return 0, err
	}
	return category.ID, nil
}
