package handler

import(
	"github.com/labstack/echo/v4"
)

type userHandler struct{}

func newUserHandler() *userHandler{
	return &userHandler{}
}

func (h *userHandler) home(c echo.Context) error{
	return c.JSON(200, echo.Map{
		"message": "user home",
	})
}