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

func (r *categoryRepository) GetCategoryTasks(id uint) (model.Category, error) {
	var category model.Category
	err := r.db.
		Preload("Tasks").
		First(&category, id).
		Error
	if err != nil {
		return model.Category{}, err
	}
	return category, nil
}

func (r *categoryRepository) CreateTask(task *model.Task) (uint, error) {
	if err := r.db.Create(&task).Error; err != nil {
		return 0, err
	}
	return task.ID, nil
}

func (r *categoryRepository) DeleteTask(taskID uint) error {
	err := r.db.Delete(&model.Task{}, taskID).Error
	return err
}
