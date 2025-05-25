package handler

import(
	"github.com/labstack/echo/v4"
)

type adminHandler struct{}

func newAdminHandler() *adminHandler{
	return &adminHandler{}
}

func (h *adminHandler) home(c echo.Context) error{
	return c.JSON(200, echo.Map{
		"message": "admin home",
	})
}