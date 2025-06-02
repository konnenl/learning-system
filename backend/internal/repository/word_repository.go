package repository

import (
	"gorm.io/gorm"
	"github.com/konnenl/learning-system/internal/model"
)

type wordRepository struct {
	db *gorm.DB
}

func newWordRepository(db *gorm.DB) *wordRepository {
	return &wordRepository{db: db}
}

func (r *wordRepository) GetWords() ([]*model.Word, error) {
	var words []*model.Word
	level := []string{"A1", "A2", "B1", "B2", "C1", "C2"}

	for _, lev := range level{
		temp_words := []*model.Word{}
		err := r.db.
			Where("level = ?", lev).
			Order("RANDOM()").
			Limit(3).
			Find(&temp_words).
			Error
		if err != nil{
			return nil, err
		}
		words = append(words, temp_words...)
	}
	return words, nil
}