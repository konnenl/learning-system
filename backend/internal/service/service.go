package service

import(
	"github.com/labstack/echo/v4"
	"github.com/konnenl/learning-system/internal/repository"
)

type Service struct{
	Model ModelService
	Lesson LessonService
	Auth AuthService
}

func NewService(repository *repository.Repository, key string, expires int) *Service{
	return &Service{
		Model: newModelService(),
		Lesson: newLessonService(),
		Auth: newJWTService(key, expires),
	}
}

type ModelService interface{
	GetLevel() string
}

type LessonService interface{
}

type AuthService interface{
	GenerateToken(userId uint, role string) (string, error)
	Middleware() echo.MiddlewareFunc
	GetClaims(c echo.Context) (*Claims, error)
}