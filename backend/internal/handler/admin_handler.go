package handler

import(
	_"github.com/labstack/echo/v4"
)

type adminHandler struct{}

func newAdminHandler() *adminHandler{
	return &adminHandler{}
}
