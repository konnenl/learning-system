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
	e.POST("/login", h.auth.loginPost)
	e.GET("/register", h.auth.register)
	e.POST("/register", h.auth.registerPost)
	e.POST("/logout", h.auth.logoutPost)

	user := e.Group("/user")
	user.GET("/", h.user.home)
	user.GET("/test", h.user.test)
	user.POST("/test", h.user.testPost)

	admin := e.Group("/admin")
	admin.GET("/home", h.user.home)
	admin.GET("/create-topic", h.user.home)
	admin.GET("/create-question", h.user.home)
	admin.GET("/topics", h.user.home)
}
