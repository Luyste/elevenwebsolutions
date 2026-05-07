package render

import (
	"html/template"
	"io"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func NewTemplate() *Template {
	files, _ := filepath.Glob("web/views/*.html")
	partials, _ := filepath.Glob("web/views/partials/*.html")
	files = append(files, partials...)

	return &Template{
		templates: template.Must(template.ParseFiles(files...)),
	}
}

func (t *Template) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
