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
	return c.Render(200, "auth/login", map[string]interface{}{
		"Title": "Авторизация",
	})
}

func (h *authHandler) loginPost(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"message": "ok",
	})
}

func (h *authHandler) logoutPost(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"message": "ok",
	})
}

func (h *authHandler) register(c echo.Context) error {
	return c.Render(200, "auth/register", map[string]interface{}{
		"Title": "Register",
	})
}

func (h *authHandler) registerPost(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"message": "ok",
	})
}

