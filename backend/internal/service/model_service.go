package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/konnenl/learning-system/internal/repository"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type modelService struct {
	wordRepository repository.WordRepository
	userRepository repository.UserRepository
	httpClient     *http.Client
}

func newModelService(wordRepository repository.WordRepository, userRepository repository.UserRepository) *modelService {
	return &modelService{
		wordRepository: wordRepository,
		userRepository: userRepository,
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (s *modelService) ProcessPlacementTest(id uint, data []WordAnswer) (string, error) {
	words_id := make([]uint, len(data))
	for i, a := range data {
		words_id[i] = a.ID
	}

	levels, err := s.wordRepository.GetLevels(words_id)
	if err != nil {
		return "", err
	}

	know_by_level := map[string]float32{
		"A1": 0,
		"A2": 0,
		"B1": 0,
		"B2": 0,
		"C1": 0,
		"C2": 0,
	}
	for i, lev := range levels {
		if data[i].Know {
			know_by_level[lev]++
		}
	}

	x := []float32{}
	n := float32(len(data) / 6)
	for _, lev := range know_by_level {
		x = append(x, lev/n)
	}
	level, err := s.PredictLevel(x)
	if err != nil {
		return "", err
	}

	err = s.userRepository.UpdateLevel(id, level)
	if err != nil {
		return "", err
	}
	return level, nil
}

func (s *modelService) PredictLevel(x []float32) (string, error) {
	jsonBody, err := json.Marshal(echo.Map{
		"x": x,
	})
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}
	res, err := s.httpClient.Post(
		"http://localhost:5000/predict",
		"application/json",
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return "", fmt.Errorf("API request failed: %w", err)
	}
	defer res.Body.Close()

	var apiResponse struct {
		Status       string    `json:"status"`
		Result       string    `json:"result"`
		ReceivedData []float32 `json:"received_data"`
	}

	if err := json.NewDecoder(res.Body).Decode(&apiResponse); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}
	return apiResponse.Result, nil
}
