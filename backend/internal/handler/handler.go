package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/konnenl/learning-system/internal/service"
)

type Handler struct {
	user  *userHandler
	admin *adminHandler
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		user:  newUserHandler(service.User, service.Lesson),
		admin: newAdminHandler(),
	}
}
func (h *Handler) InitRoutes(e *echo.Echo) {
	//TODO auth middleware
	e.GET("/", h.user.home)
	e.GET("/login", h.user.login)
	e.GET("/register", h.user.register)
	e.GET("/logout", h.user.logout)
	e.GET("/test", h.user.test)

	admin := e.Group("/admin")
	admin.GET("/home", h.user.home)
	admin.GET("/create-topic", h.user.home)
	admin.GET("/create-question", h.user.home)
	admin.GET("/topics", h.user.home)
}
