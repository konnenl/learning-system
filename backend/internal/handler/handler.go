package handler

import (
	"github.com/konnenl/learning-system/internal/repository"
	"github.com/konnenl/learning-system/internal/service"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	user        *userHandler
	admin       *adminHandler
	auth        *authHandler
	authService service.AuthService
}

func NewHandler(service *service.Service, repository *repository.Repository) *Handler {
	return &Handler{
		user:        newUserHandler(service.Auth, repository.Word, repository.User, service.Model),
		admin:       newAdminHandler(service.Auth, repository.Category),
		auth:        newAuthHandler(service.Auth, repository.User),
		authService: service.Auth,
	}
}
func (h *Handler) InitRoutes(e *echo.Echo) {
	auth := e.Group("/auth")
	auth.POST("/register", h.auth.register)
	auth.POST("/login", h.auth.login)

	users := e.Group("/users")
	users.Use(h.authService.Middleware())
	users.GET("/level", h.user.getLevel)
	//users.GET("/test/next", h.user.getTest)
	//users.POST("/test/submit", h.user.submitTest)
	users.GET("/placement", h.user.getPlacementTest)
	users.POST("/placement", h.user.submitPlacementTest)

	admin := e.Group("/admin")
	admin.Use(h.authService.Middleware())
	admin.Use(h.authService.AdminMiddleware())
	admin.GET("/categories", h.admin.getAllCategories)
	admin.POST("/categories", h.admin.createCategory)
	admin.GET("/categories/:categoryID/tasks", h.admin.getTasksByCategory)
	admin.POST("/categories/:categoryID/tasks", h.admin.createTask)
	admin.DELETE("/categories/:categoryID/tasks/:taskID", h.admin.deleteTask)
	admin.POST("/users", h.admin.createAdminUser)
}
