package main

import (
	"go-and-htmx/internal/handlers"
	render "go-and-htmx/tools"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Logger.SetLevel(log.DEBUG)

	e.Static("/static", "web/static")

	e.Renderer = render.NewTemplate()

	e.GET("/", handlers.Home)
	e.POST("/contact", handlers.Contact)

	e.Logger.Fatal(e.Start(":3000"))
}
