package tl

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

// template wrapper
type tlw struct {
	Tl *template.Template
}

func (t tlw) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Tl.ExecuteTemplate(w, name, data)
}
