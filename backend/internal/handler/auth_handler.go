package handler

import (
	"github.com/konnenl/learning-system/internal/model"
	"github.com/konnenl/learning-system/internal/repository"
	"github.com/konnenl/learning-system/internal/service"
	"github.com/konnenl/learning-system/internal/validator"
	"github.com/labstack/echo/v4"
	"strings"
)

type authHandler struct {
	authService    service.AuthService
	userRepository repository.UserRepository
}

func newAuthHandler(authService service.AuthService, userRepository repository.UserRepository) *authHandler {
	return &authHandler{
		authService:    authService,
		userRepository: userRepository,
	}
}

func (h *authHandler) login(c echo.Context) error {
	var r userLoginRequest
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
	user, err := h.userRepository.Authenticate(r.Email, r.Password)
	if err != nil {
		return c.JSON(401, echo.Map{
			"error": "Invalid credentials",
		})
	}

	token, err := h.authService.GenerateToken(user.ID, user.Role)
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": "Internal error",
		})
	}
	return c.JSON(200, echo.Map{
		"id":    user.ID,
		"token": token,
	})
}

func (h *authHandler) register(c echo.Context) error {
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
		Role:     "user",
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

	token, err := h.authService.GenerateToken(id, user.Role)
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": "Internal error",
		})
	}

	return c.JSON(201, echo.Map{
		"id":    id,
		"token": token,
	})
}
