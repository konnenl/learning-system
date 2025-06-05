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
	//TODO auth middleware
	e.POST("/login", h.auth.login)
	e.POST("/register", h.auth.register)

	users := e.Group("/users")
	users.Use(h.authService.Middleware())
	users.GET("/level", h.user.getLevel) // уровень
	//users.GET("/test/next", h.user.getTest) // тест (название, вопросы)
	//users.POST("/test/submit", h.user.submitTest) // отправка ответов -> взять ответы, посчитать правильные, поменять значение progress
	users.GET("/placement", h.user.getPlacementTest)     // входной тест
	users.POST("/placement", h.user.submitPlacementTest) // отправка ответов -> взять ответы на входной тест, отправить в модель, записать уровень в бд

	admin := e.Group("/admin")
	admin.Use(h.authService.Middleware())
	admin.Use(h.authService.AdminMiddleware())
	admin.GET("/categories", h.admin.getAllCategories)
	// admin.GET("/tasks/:categoryID", h.admin.getTasksByCategory)
	admin.POST("/categories", h.admin.createCategory)
	// admin.POST("/tasks", h.admin.createTask)
	// admin.POST("/admins", h.admin.createAdminAccount)
}
