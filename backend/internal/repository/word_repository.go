package repository

import (
	"github.com/konnenl/learning-system/internal/model"
	"gorm.io/gorm"
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

	for _, lev := range level {
		temp_words := []*model.Word{}
		err := r.db.
			Where("level = ?", lev).
			Order("RANDOM()").
			Limit(3).
			Find(&temp_words).
			Error
		if err != nil {
			return nil, err
		}
		words = append(words, temp_words...)
	}
	return words, nil
}

func (r *wordRepository) GetLevels(words_id []uint) ([]string, error) {
	levels := make([]string, len(words_id))
	var temp string
	for i, id := range words_id {
		err := r.db.
			Model(&model.Word{}).
			Where("id = ?", id).
			Select("level").
			Scan(&temp).
			Error
		if err != nil {
			return nil, err
		}
		levels[i] = temp
	}
	return levels, nil
}
