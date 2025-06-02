package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/konnenl/learning-system/internal/service"
)

type Handler struct {
	user  *userHandler
	admin *adminHandler
	auth *authHandler
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		user:  newUserHandler(service.User, service.Lesson),
		admin: newAdminHandler(),
		auth: newAuthHandler(service.User),
	}
}
func (h *Handler) InitRoutes(e *echo.Echo) {
	//TODO auth middleware
	e.GET("/login", h.auth.login)
	e.GET("/register", h.auth.register)
	e.POST("/logout", h.auth.logout)

	user := e.Group("/user")
	user.GET("/", h.user.home)
	user.GET("/test", h.user.test)

	//admin := e.Group("/admin")
}
