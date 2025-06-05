package handler

import (
	"github.com/konnenl/learning-system/internal/repository"
	"github.com/konnenl/learning-system/internal/service"
	"github.com/konnenl/learning-system/internal/validator"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	authService    service.AuthService
	modelService   service.ModelService
	wordRepository repository.WordRepository
	userRepository repository.UserRepository
}

func newUserHandler(authService service.AuthService, wordRepository repository.WordRepository, userRepository repository.UserRepository, modelService service.ModelService) *userHandler {
	return &userHandler{
		authService:    authService,
		modelService:   modelService,
		wordRepository: wordRepository,
		userRepository: userRepository,
	}
}

// users.GET("/level", h.user.getLevel) // уровень
func (h *userHandler) getLevel(c echo.Context) error {
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

	return c.JSON(200, echo.Map{
		"message": "ok",
		"level":   level,
	})
}

// users.GET("/test/next", h.user.getTest) // тест (название, вопросы)
// users.POST("/test/submit", h.user.submitTest) // отправка ответов -> взять ответы, посчитать правильные, поменять значение progress

// users.GET("/placement", h.user.getPlacementTest) // входной тест
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

// users.POST("/placement", h.user.submitPlacementTest) // отправка ответов -> взять ответы на входной тест, отправить в модель, записать уровень в бд
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
