package handler

import(
	"github.com/konnenl/learning-system/internal/model"
)

type wordResponce struct{
	ID uint `json:"id"`
	Word string `json:"word"`
}

func newPlacementTestResponce(w []*model.Word) []wordResponce {
	words := make([]wordResponce, len(w))
	for i, word := range w {
		word_responce := wordResponce{
			ID: word.ID,
			Word: word.Word,
		}
		words[i] = word_responce
	}

	return words
}