package service

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/konnenl/learning-system/internal/repository"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"strings"
	"time"
)

type JWTService struct {
	SecretKey      []byte
	Expires        int
	userRepository repository.UserRepository
}

type Claims struct {
	UserId uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func newJWTService(key string, expires int, userRepository repository.UserRepository) *JWTService {
	return &JWTService{
		SecretKey:      []byte(key),
		Expires:        expires,
		userRepository: userRepository,
	}
}

func (s *JWTService) GenerateToken(userId uint, role string) (string, error) {
	expiresAt := time.Now().Add(time.Duration(s.Expires) * time.Second)
	claims := &Claims{
		UserId: userId,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.SecretKey)
}

func (s *JWTService) Middleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    s.SecretKey,
		SigningMethod: "HS256",
		TokenLookup:   "header:Authorization",
		ContextKey:    "user",
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return &Claims{}
		},
		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
			parts := strings.Split(auth, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return nil, echo.NewHTTPError(401, "Invalid authorization header format")
			}
			token, err := jwt.ParseWithClaims(parts[1], &Claims{}, func(token *jwt.Token) (interface{}, error) {
				return s.SecretKey, nil
			})
			if err != nil {
				return nil, echo.NewHTTPError(401, "Invalid token")
			}

			if claims, ok := token.Claims.(*Claims); ok {
				_, err := s.userRepository.GetByID(claims.UserId)
				if err != nil {
					return nil, echo.NewHTTPError(401, "User not found")
				}
			}

			return token, nil
		},
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(401, echo.Map{
				"error": "Unauthorized",
			})
		},
	})
}

func (s *JWTService) AdminMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims, err := s.GetClaims(c)
			if err != nil {
				return err
			}

			if claims.Role != "admin" {
				return echo.NewHTTPError(403, "Forbidden")
			}
			return next(c)
		}
	}
}

func (s *JWTService) UserMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims, err := s.GetClaims(c)
			if err != nil {
				return err
			}

			if claims.Role != "user" {
				return echo.NewHTTPError(403, "Forbidden")
			}
			return next(c)
		}
	}
}

func (s *JWTService) GetClaims(c echo.Context) (*Claims, error) {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return nil, echo.NewHTTPError(401, "invalid token")
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, echo.NewHTTPError(401, "invalid claims")
	}
	if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, echo.NewHTTPError(401, "token expired")
	}
	return claims, nil
}
