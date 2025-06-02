package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/konnenl/learning-system/internal/service"
	"github.com/konnenl/learning-system/internal/repository"
)

type Handler struct {
	user  *userHandler
	admin *adminHandler
	auth *authHandler
	authService service.AuthService
}

func NewHandler(service *service.Service, repository *repository.Repository) *Handler {
	return &Handler{
		user:  newUserHandler(service.Auth),
		admin: newAdminHandler(service.Auth),
		auth: newAuthHandler(service.Auth, repository.User),
		authService: service.Auth,
	}
}
func (h *Handler) InitRoutes(e *echo.Echo) {
	//TODO auth middleware
	e.POST("/login", h.auth.login)
	e.POST("/register", h.auth.register)

	users := e.Group("/user")
	users.Use(h.authService.Middleware())
	users.GET("/", h.user.home)
	users.GET("/test", h.user.test)
	users.POST("/predict", h.user.predict)

	//admin := e.Group("/admin")
	//admin.Use(h.authService.Middleware())
	//admin.Use(h.authService.AdminMiddleware())
}
