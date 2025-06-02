package handler

import (
	_ "github.com/labstack/echo/v4"
	"github.com/konnenl/learning-system/internal/service"
)

type adminHandler struct{
	authService service.AuthService
}

func newAdminHandler(authService service.AuthService) *adminHandler {
	return &adminHandler{
		authService: authService,
	}
}
