package handler

import (
	"github.com/konnenl/learning-system/internal/model"
)

type wordResponce struct {
	ID   uint   `json:"id"`
	Word string `json:"word"`
}

func newPlacementTestResponce(w []*model.Word) []wordResponce {
	words := make([]wordResponce, len(w))
	for i, word := range w {
		word_responce := wordResponce{
			ID:   word.ID,
			Word: word.Word,
		}
		words[i] = word_responce
	}

	return words
}

type categoryResponce struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Level string `json:"level"`
}

func newCategoriesResponce(c []model.Category) []categoryResponce {
	categories := make([]categoryResponce, len(c))
	for i, category := range c {
		category_responce := categoryResponce{
			ID:    category.ID,
			Name:  category.Name,
			Level: category.Level,
		}
		categories[i] = category_responce
	}

	return categories
}
