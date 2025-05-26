package template

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func InitTemplate() *Template {
	templates := template.New("")

	template.Must(templates.ParseGlob("ui/html/partials/*.html"))
	template.Must(templates.ParseGlob("ui/html/pages/user/*.html"))
	//template.Must(templates.ParseGlob("ui/html/pages/admin/*.html"))

	return &Template{
		templates: templates,
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	for _, tmpl := range t.templates.Templates() {
		fmt.Println("Available template:", tmpl.Name())
	}
	return t.templates.ExecuteTemplate(w, name, data)
}
