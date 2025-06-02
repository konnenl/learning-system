package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/konnenl/learning-system/internal/service"
)

type userHandler struct{
	authService service.AuthService
}

func newUserHandler(authService service.AuthService) *userHandler {
	return &userHandler{
		authService: authService,
	}
}

func (h *userHandler) home(c echo.Context) error {
	level := "A1"
	testTitle := "Тест по теме животные"
	return c.JSON(200,map[string]interface{}{
		"HasLevel": true,
		"Level": level,
		"TestTitle": testTitle,
	})
}

func (h *userHandler) test(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"message": "ok",
	})
}

func (h *userHandler) predict(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"message": "ok",
	})
}