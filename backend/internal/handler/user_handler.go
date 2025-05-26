package handler

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/konnenl/learning-system/internal/service"
)

type userHandler struct{
	userService service.UserService
	lessonService service.LessonService
}

func newUserHandler(userService service.UserService, lessonService service.LessonService) *userHandler {
	return &userHandler{
		userService: userService,
		lessonService: lessonService,
	}
}

func (h *userHandler) home(c echo.Context) error {
	//TODO логика получения данны[]
	level := "A1"
	testTitle := "Тест по теме животные"
	return c.Render(200, "user/home", map[string]interface{}{
		"HasLevel": true,
		"Level": level,
		"TestTitle": testTitle,
	})
}

func (h *userHandler) login(c echo.Context) error {
	if c.Request().Method == http.MethodPost {
		//email := c.FormValue("email")
		//password := c.FormValue("password")
		
		// TODO логика авторизации
		return c.Redirect(302, "/")
	}

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
	return c.Render(200, "user/register", map[string]interface{}{
		"Title": "Register",
	})
}

func (h *userHandler) test(c echo.Context) error {
	return c.Render(200, "user/test", map[string]interface{}{
		"Title": "Прохождение теста",
	})
}
