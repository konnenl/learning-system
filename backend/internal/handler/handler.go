package handler

import(
	"github.com/labstack/echo/v4"
)

type Handler struct{
	user *userHandler
	admin *adminHandler
}

func NewHandler() *Handler {
	return &Handler{
		user: newUserHandler(),
		admin: newAdminHandler(),
	}
}
func (h *Handler) InitRoutes(e *echo.Echo) {
	e.GET("/", h.user.home)
	e.GET("/admin", h.admin.home)
}