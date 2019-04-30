package main

import (
	"html/template"
	"io"
	"log"
	"os"

	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

//Specification for the service
type Specification struct {
	ServerAddr string `env:"SERVER_ADDR"`
}

type App struct {
	Spec *Specification
}

type Template struct {
	templates *template.Template
}

//Render template handler
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	log.Printf("site: starting")
	s := &Specification{
		ServerAddr: getenv("SERVER_ADDR", "0.0.0.0:7111"),
	}

	app := App{Spec: s}

	t := &Template{
		templates: template.Must(template.ParseGlob("static/*.html")),
	}

	e := echo.New()

	e.Renderer = t
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.GET("/", app.HomeHandler)

	e.Logger.Fatal(e.Start(s.ServerAddr))
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
