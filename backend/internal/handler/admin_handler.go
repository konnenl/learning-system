package handler

import (
	"github.com/konnenl/learning-system/internal/model"
	"github.com/konnenl/learning-system/internal/repository"
	"github.com/konnenl/learning-system/internal/service"
	"github.com/konnenl/learning-system/internal/validator"
	"github.com/labstack/echo/v4"
	"strconv"
	"strings"
)

type adminHandler struct {
	authService        service.AuthService
	categoryRepository repository.CategoryRepository
	userRepository     repository.UserRepository
	wordRepository     repository.WordRepository
}

func newAdminHandler(authService service.AuthService, categoryRepository repository.CategoryRepository, userRepository repository.UserRepository, wordRepository repository.WordRepository) *adminHandler {
	return &adminHandler{
		authService:        authService,
		categoryRepository: categoryRepository,
		userRepository:     userRepository,
		wordRepository:     wordRepository,
	}
}

func (h *adminHandler) getAllCategories(c echo.Context) error {
	categories, err := h.categoryRepository.GetAllCategories()
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": "Internal error",
		})
	}
	categories_responce := newCategoriesResponce(categories)
	return c.JSON(200, echo.Map{
		"categories": categories_responce,
	})
}

func (h *adminHandler) createCategory(c echo.Context) error {
	var r categoryRequest
	if err := c.Bind(&r); err != nil {
		return c.JSON(400, echo.Map{
			"error": "Bad request",
		})
	}

	if err := c.Validate(r); err != nil {
		return c.JSON(400, echo.Map{
			"error":  "Validation failed",
			"fields": validator.GetValidationErrors(err),
		})
	}

	category := &model.Category{
		Name:  r.Name,
		Level: r.Level,
	}

	id, err := h.categoryRepository.Create(category)
	if err != nil {
		return c.JSON(400, echo.Map{
			"error": "Failed to create category",
		})
	}

	return c.JSON(200, echo.Map{
		"id": id,
	})
}

func (h *adminHandler) deleteCategory(c echo.Context) error {
	categoryID, err := strconv.Atoi(c.Param("categoryID"))
	if err != nil {
		return c.JSON(400, echo.Map{"error": "Invalid categoryID"})
	}

	err = h.categoryRepository.Delete(uint(categoryID))
	if err != nil {
		return c.JSON(400, echo.Map{
			"error": "Failed to delete category",
		})
	}

	return c.JSON(200, echo.Map{
		"message": "Category deleted",
	})
}

func (h *adminHandler) getTasksByCategory(c echo.Context) error {
	categoryID, err := strconv.Atoi(c.Param("categoryID"))
	if err != nil {
		return c.JSON(400, echo.Map{"error": "Invalid categoryID"})
	}

	category, err := h.categoryRepository.GetCategoryTasks(uint(categoryID))
	if err != nil {
		return c.JSON(400, echo.Map{
			"error": "Failed to get category tasks",
		})
	}
	categoryTasksResponce := newCategoryTasksResponce(category)
	return c.JSON(200, echo.Map{
		"category_tasks": categoryTasksResponce,
	})
}

func (h *adminHandler) createTask(c echo.Context) error {
	categoryID, err := strconv.Atoi(c.Param("categoryID"))
	if err != nil {
		return c.JSON(400, echo.Map{"error": "Invalid categoryID"})
	}

	var r taskRequest
	if err := c.Bind(&r); err != nil {
		return c.JSON(400, echo.Map{
			"error": "Bad request",
		})
	}

	if err := c.Validate(r); err != nil {
		return c.JSON(400, echo.Map{
			"error":  "Validation failed",
			"fields": validator.GetValidationErrors(err),
		})
	}

	task := &model.Task{
		CategoryID:  uint(categoryID),
		Description: r.Description,
		Question:    r.Question,
		Answer:      r.Answer,
	}

	id, err := h.categoryRepository.CreateTask(task)
	if err != nil {
		return c.JSON(400, echo.Map{
			"error": "Failed to create task",
		})
	}

	return c.JSON(200, echo.Map{
		"id": id,
	})
}

func (h *adminHandler) deleteTask(c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("taskID"))
	if err != nil {
		return c.JSON(400, echo.Map{"error": "Invalid taskID"})
	}

	err = h.categoryRepository.DeleteTask(uint(taskID))
	if err != nil {
		return c.JSON(400, echo.Map{
			"error": "Failed to delete task",
		})
	}

	return c.JSON(200, echo.Map{
		"message": "Task deleted",
	})
}

func (h *adminHandler) createAdminUser(c echo.Context) error {
	var r userRegisterRequest
	if err := c.Bind(&r); err != nil {
		return c.JSON(400, echo.Map{
			"error": "Bad request",
		})
	}

	if err := c.Validate(r); err != nil {
		return c.JSON(400, echo.Map{
			"error":  "Validation failed",
			"fields": validator.GetValidationErrors(err),
		})
	}
	hashedPassword, err := model.HashPassword(r.Password)
	if err != nil {
		return c.JSON(400, echo.Map{
			"error": "Bad request",
		})
	}

	user := &model.User{
		Fullname: r.Fullname,
		Email:    r.Email,
		Password: hashedPassword,
		Role:     "admin",
	}

	id, err := h.userRepository.Create(user)
	if err != nil {
		if strings.Contains(err.Error(), "email already exist") {
			return c.JSON(409, echo.Map{
				"error": "Email already in use",
			})
		}
		return c.JSON(500, echo.Map{
			"error": "Failed to create user",
		})
	}

	return c.JSON(201, echo.Map{
		"id": id,
	})
}

func (h *adminHandler) getAllAdminUser(c echo.Context) error {
	admins, err := h.userRepository.GetAllAdmin()
	if err != nil {
		return c.JSON(400, echo.Map{
			"error": "Internal error",
		})
	}
	adminsResponce := newUsersResponce(admins)
	return c.JSON(200, echo.Map{
		"admins": adminsResponce,
	})
}

func (h *adminHandler) deleteAdminUser(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		return c.JSON(400, echo.Map{"error": "Invalid userID"})
	}

	err = h.userRepository.Delete(uint(userID))
	if err != nil {
		return c.JSON(400, echo.Map{
			"error": "Failed to delete user",
		})
	}

	return c.JSON(200, echo.Map{
		"message": "Users deleted",
	})
}

func (h *adminHandler) getAllWords(c echo.Context) error {
	words, err := h.wordRepository.GetAllWords()
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": "Internal error",
		})
	}
	words_responce := newWordsResponce(words)
	return c.JSON(200, echo.Map{
		"words": words_responce,
	})
}

func (h *adminHandler) createWord(c echo.Context) error {
	var r wordRequest
	if err := c.Bind(&r); err != nil {
		return c.JSON(400, echo.Map{
			"error": "Bad request",
		})
	}

	if err := c.Validate(r); err != nil {
		return c.JSON(400, echo.Map{
			"error":  "Validation failed",
			"fields": validator.GetValidationErrors(err),
		})
	}

	word := &model.Word{
		Word:  r.Word,
		Level: r.Level,
	}

	id, err := h.wordRepository.Create(word)
	if err != nil {
		return c.JSON(400, echo.Map{
			"error": "Failed to create word",
		})
	}

	return c.JSON(200, echo.Map{
		"id": id,
	})
}

func (h *adminHandler) deleteWord(c echo.Context) error {
	wordID, err := strconv.Atoi(c.Param("wordID"))
	if err != nil {
		return c.JSON(400, echo.Map{"error": "Invalid wordID"})
	}

	err = h.wordRepository.Delete(uint(wordID))
	if err != nil {
		return c.JSON(400, echo.Map{
			"error": "Failed to delete word",
		})
	}

	return c.JSON(200, echo.Map{
		"message": "Word deleted",
	})
}
