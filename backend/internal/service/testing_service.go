package service

import (
	"github.com/konnenl/learning-system/internal/repository"
	"strings"
)

type testingService struct {
	categoryRepository repository.CategoryRepository
	userRepository     repository.UserRepository
}

func newTestingService(categoryRepository repository.CategoryRepository, userRepository repository.UserRepository) *testingService {
	return &testingService{
		categoryRepository: categoryRepository,
		userRepository:     userRepository,
	}
}

func (s *testingService) ProcessTest(userID uint, categoryID uint, data []TaskAnswer) (uint, error) {
	category, _ := s.categoryRepository.GetCategoryTasks(categoryID)
	correctAnswers := make(map[uint]string)
	for _, task := range category.Tasks {
		correctAnswers[task.ID] = task.Answer
	}

	correctCount := 0
	for _, answer := range data {
		correctAnswer, exists := correctAnswers[answer.TaskID]
		if exists && strings.EqualFold(strings.ToLower(strings.TrimSpace(answer.Answer)), strings.ToLower(strings.TrimSpace(correctAnswer))) {
			correctCount++
		}
	}

	progress := float32(correctCount) / float32(len(data))

	err := s.categoryRepository.UpdateProgress(userID, categoryID, progress)
	if err != nil {
		return 0, err
	}

	return uint(correctCount), nil
}
