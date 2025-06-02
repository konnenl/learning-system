package service

import (
	"github.com/konnenl/learning-system/internal/repository"
	"github.com/labstack/echo/v4"
)

type Service struct {
	Model ModelService
	Auth  AuthService
}

func NewService(repository *repository.Repository, key string, expires int) *Service {
	return &Service{
		Model: newModelService(repository.Word, repository.User),
		Auth:  newJWTService(key, expires),
	}
}

type ModelService interface {
	ProcessPlacementTest(id uint, data []WordAnswer) (string, error)
}

type AuthService interface {
	GenerateToken(userId uint, role string) (string, error)
	Middleware() echo.MiddlewareFunc
	GetClaims(c echo.Context) (*Claims, error)
}
