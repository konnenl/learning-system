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

type taskResponce struct {
	ID          uint   `json:"id"`
	Description string `json:"description"`
	Question    string `json:"question"`
	Answer      string `json:"answer"`
}

type categoryTasksResponce struct {
	Category categoryResponce `json:"category"`
	Tasks    []taskResponce   `json:"tasks"`
}

func NewCategoryTasksResponce(c model.Category) categoryTasksResponce {
	categoryTasks := categoryTasksResponce{
		Category: categoryResponce{
			ID:    c.ID,
			Name:  c.Name,
			Level: c.Level,
		},
	}
	tasks := make([]taskResponce, len(c.Tasks))
	for i, task := range c.Tasks {
		task := taskResponce{
			ID:          task.ID,
			Description: task.Description,
			Question:    task.Question,
			Answer:      task.Answer,
		}
		tasks[i] = task
	}
	categoryTasks.Tasks = tasks
	return categoryTasks
}

func NewTestResponce(c model.Category) categoryTasksResponce {
	categoryTasks := categoryTasksResponce{
		Category: categoryResponce{
			ID:    c.ID,
			Name:  c.Name,
			Level: c.Level,
		},
	}
	tasks := make([]taskResponce, len(c.Tasks))
	for i, task := range c.Tasks {
		task := taskResponce{
			ID:          task.ID,
			Description: task.Description,
			Question:    task.Question,
		}
		tasks[i] = task
	}
	categoryTasks.Tasks = tasks
	return categoryTasks
}

type userResponce struct {
	ID       uint   `jsnon:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

func newUsersResponce(u []model.User) []userResponce {
	users := make([]userResponce, len(u))
	for i, user := range u {
		userResponce := userResponce{
			ID:       user.ID,
			Fullname: user.Fullname,
			Email:    user.Email,
		}
		users[i] = userResponce
	}
	return users
}
func newUserResponce(u *model.User) userResponce {
	userResponce := userResponce{
		ID:       u.ID,
		Fullname: u.Fullname,
		Email:    u.Email,
	}
	return userResponce
}

type wordAdminResponce struct {
	ID    uint   `json:"id"`
	Word  string `json:"word"`
	Level string `json:"level"`
}

func newWordsResponce(w []model.Word) []wordAdminResponce {
	words := make([]wordAdminResponce, len(w))
	for i, word := range w {
		wordResponce := wordAdminResponce{
			ID:    word.ID,
			Word:  word.Word,
			Level: word.Level,
		}
		words[i] = wordResponce
	}
	return words
}
