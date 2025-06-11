package handler

import (
	"github.com/konnenl/learning-system/internal/repository"
	"github.com/konnenl/learning-system/internal/service"
	"github.com/konnenl/learning-system/internal/validator"
	"github.com/labstack/echo/v4"
	"strconv"
)

type userHandler struct {
	authService        service.AuthService
	modelService       service.ModelService
	testingService     service.TestingService
	wordRepository     repository.WordRepository
	userRepository     repository.UserRepository
	categoryRepository repository.CategoryRepository
}

func newUserHandler(authService service.AuthService,
	wordRepository repository.WordRepository,
	userRepository repository.UserRepository,
	categoryRepository repository.CategoryRepository,
	modelService service.ModelService,
	testingService service.TestingService) *userHandler {
	return &userHandler{
		authService:        authService,
		modelService:       modelService,
		testingService:     testingService,
		wordRepository:     wordRepository,
		userRepository:     userRepository,
		categoryRepository: categoryRepository,
	}
}

// user.GET("", h.user.getUser)
func (h *userHandler) getUser(c echo.Context) error {
	claims, err := h.authService.GetClaims(c)
	if err != nil {
		if httpErr, ok := err.(*echo.HTTPError); ok {
			return httpErr
		}
		return echo.NewHTTPError(401, "Invalid authentication")
	}

	id := uint(claims.UserId)

	user, err := h.userRepository.GetByID(id)
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": "Failed to get user",
		})
	}
	userResponce := newUserResponce(user)
	return c.JSON(200, echo.Map{
		"user": userResponce,
	})
}

//users.DELETE(":userID", h.user.deleteUser)
func (h *userHandler) deleteUser(c echo.Context) error {
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

// users.GET("/level", h.user.getLevel)
func (h *userHandler) getLevel(c echo.Context) error {
	claims, err := h.authService.GetClaims(c)
	if err != nil {
		if httpErr, ok := err.(*echo.HTTPError); ok {
			return httpErr
		}
		return echo.NewHTTPError(401, "Invalid authentication")
	}

	id := uint(claims.UserId)

	category, err := h.categoryRepository.GetNextCategory(id)

	if err != nil {
		return c.JSON(500, echo.Map{
			"error": "Internal error",
		})
	}

	if category.ID == 0 {
		level, err := h.userRepository.GetLevel(id)
		if err != nil {
			return c.JSON(500, echo.Map{
				"error": "Internal error",
			})
		}
		if level == "C2" {
			return c.JSON(200, echo.Map{
				"message": "You've completed all available English levels and all tasks",
				"level":   level,
			})
		}
		level, err = h.userRepository.UpgrateLevel(id)
		if err != nil {
			return c.JSON(500, echo.Map{
				"error": "Internal error",
			})
		}
		return c.JSON(200, echo.Map{
			"message": "Level upgraded",
			"level":   level,
		})
	}

	level, err := h.userRepository.GetLevel(id)

	if err != nil {
		return c.JSON(500, echo.Map{
			"error": "Internal error",
		})
	}

	return c.JSON(200, echo.Map{
		"level": level,
	})
}

// users.GET("/test/next", h.user.getTest)
func (h *userHandler) getTest(c echo.Context) error {
	claims, err := h.authService.GetClaims(c)
	if err != nil {
		if httpErr, ok := err.(*echo.HTTPError); ok {
			return httpErr
		}
		return echo.NewHTTPError(401, "Invalid authentication")
	}

	id := uint(claims.UserId)

	level, err := h.userRepository.GetLevel(id)
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": "Internal error",
		})
	}
	if level == "" {
		return c.JSON(200, echo.Map{
			"error": "Need to take placement test first",
		})
	}

	category, err := h.categoryRepository.GetNextCategory(id)

	if err != nil {
		return c.JSON(500, echo.Map{
			"error": "Internal error",
		})
	}
	if category.ID == 0 {
		level, err := h.userRepository.GetLevel(id)
		if err != nil {
			return c.JSON(500, echo.Map{
				"error": "Internal error",
			})
		}
		if level == "C2" {
			return c.JSON(200, echo.Map{
				"message": "You've completed all available English levels and all tasks",
			})
		}
		level, err = h.userRepository.UpgrateLevel(id)
		if err != nil {
			return c.JSON(500, echo.Map{
				"error": "Internal error",
			})
		}
		return c.JSON(200, echo.Map{
			"message": "Level upgraded",
			"level":   level,
		})
	}

	categoryTasksResponce := NewTestResponce(category)

	return c.JSON(200, echo.Map{
		"test": categoryTasksResponce,
	})
}

// users.POST("/test/submit", h.user.submitTest)
func (h *userHandler) submitTest(c echo.Context) error {
	claims, err := h.authService.GetClaims(c)
	if err != nil {
		if httpErr, ok := err.(*echo.HTTPError); ok {
			return httpErr
		}
		return echo.NewHTTPError(401, "Invalid authentication")
	}
	id := uint(claims.UserId)

	categoryID, err := strconv.Atoi(c.Param("categoryID"))
	if err != nil {
		return c.JSON(400, echo.Map{"error": "Invalid categoryID"})
	}

	level, err := h.userRepository.GetLevel(id)
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": "Internal error",
		})
	}
	if level == "" {
		return c.JSON(200, echo.Map{
			"error": "Need to take placement test first",
		})
	}

	var r testRequest
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

	data := make([]service.TaskAnswer, len(r.Answers))
	for i, a := range r.Answers {
		data[i] = service.TaskAnswer{
			TaskID: a.TaskID,
			Answer: a.Answer,
		}
	}

	correct, err := h.testingService.ProcessTest(id, uint(categoryID), data)
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": "Internal error",
		})
	}

	return c.JSON(200, echo.Map{
		"correct": correct,
	})
}

// users.GET("/placement", h.user.getPlacementTest)
func (h *userHandler) getPlacementTest(c echo.Context) error {
	claims, err := h.authService.GetClaims(c)
	if err != nil {
		if httpErr, ok := err.(*echo.HTTPError); ok {
			return httpErr
		}
		return echo.NewHTTPError(401, "Invalid authentication")
	}

	id := uint(claims.UserId)
	level, err := h.userRepository.GetLevel(id)

	if err != nil {
		return c.JSON(500, echo.Map{
			"error": "Internal error",
		})
	}

	if level != "" {
		return c.JSON(200, echo.Map{
			"error": "Placement test already completed",
		})
	}

	words, err := h.wordRepository.GetWords()
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": "Internal error",
		})
	}

	placement_test_responce := newPlacementTestResponce(words)
	return c.JSON(200, echo.Map{
		"message":        "Ok",
		"placement_test": placement_test_responce,
	})
}

// users.POST("/placement", h.user.submitPlacementTest)
func (h *userHandler) submitPlacementTest(c echo.Context) error {
	claims, err := h.authService.GetClaims(c)
	if err != nil {
		if httpErr, ok := err.(*echo.HTTPError); ok {
			return httpErr
		}
		return echo.NewHTTPError(401, "Invalid authentication")
	}
	id := uint(claims.UserId)

	var r placementTestRequest
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

	data := make([]service.WordAnswer, len(r.Answers))
	for i, a := range r.Answers {
		data[i] = service.WordAnswer{
			ID:   a.ID,
			Know: a.Know,
		}
	}

	level, err := h.modelService.ProcessPlacementTest(id, data)
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": "Internal error",
		})
	}

	return c.JSON(200, echo.Map{
		"message": "Ok",
		"level":   level,
	})
}
