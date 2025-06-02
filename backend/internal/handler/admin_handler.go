package handler

import (
	"github.com/konnenl/learning-system/internal/service"
	_ "github.com/labstack/echo/v4"
)

type adminHandler struct {
	authService service.AuthService
}

func newAdminHandler(authService service.AuthService) *adminHandler {
	return &adminHandler{
		authService: authService,
	}
}
