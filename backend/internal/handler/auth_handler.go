package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/konnenl/learning-system/internal/service"
)

type authHandler struct{
	userService service.UserService
}

func newAuthHandler(userService service.UserService) *authHandler {
	return &authHandler{
		userService: userService,
	}
}

func (h *authHandler) login(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"message": "ok",
	})
}

func (h *authHandler) logout(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"message": "ok",
	})
}

func (h *authHandler) register(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"message": "ok",
	})
}
