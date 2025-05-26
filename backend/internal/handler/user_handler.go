package handler

import(
	"github.com/labstack/echo/v4"
)

type userHandler struct{}

func newUserHandler() *userHandler{
	return &userHandler{}
}

func (h *userHandler) home(c echo.Context) error {
	return c.Render(200, "user/home", map[string]interface{}{
		"Title": "Главная страница",
	})
}

func (h *userHandler) login(c echo.Context) error {
	return c.Render(200, "user/login", map[string]interface{}{
		"Title": "Авторизация",
	})
}

func (h *userHandler) logout(c echo.Context) error {
	return c.Render(200, "user/login", map[string]interface{}{
		"Title": "Logout",
	})
}

func (h *userHandler) register(c echo.Context) error {
	return c.Render(200, "user/test", map[string]interface{}{
		"Title": "Register",
	})
}

func (h *userHandler) test(c echo.Context) error {
	return c.Render(200, "user/test", map[string]interface{}{
		"Title": "Прохождение теста",
	})
}