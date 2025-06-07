package repository

import (
	"github.com/konnenl/learning-system/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (r *categoryRepository) Delete(categoryID uint) error {
	err := r.db.Delete(&model.Category{}, categoryID).Error
	return err
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

func (r *categoryRepository) GetNextCategory(id uint) (model.Category, error) {
	var user model.User
	err := r.db.
		Where("id = ?", id).
		First(&user).
		Error
	if err != nil {
		return model.Category{}, err
	}

	var category model.Category
	err = r.db.
		Preload("Tasks").
		Where("level = ?", user.Level).
		Joins("LEFT JOIN progresses ON progresses.category_id = categories.id AND progresses.user_id = ?", user.ID).
		Where("progresses.value < 1 OR progresses.value IS NULL").
		Order("COALESCE(progresses.value, 0) ASC").
		Limit(1).
		Find(&category).
		Error
	if err != nil {
		return model.Category{}, err
	}
	return category, nil
}

func (r *categoryRepository) UpdateProgress(userID uint, categoryID uint, value float32) error {
	progress := &model.Progress{
		UserID:     userID,
		CategoryID: categoryID,
		Value:      value,
	}

	err := r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "category_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"value", "updated_at"}),
	}).Create(progress).Error

	return err
}
