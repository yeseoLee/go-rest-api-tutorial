package main

import (
	"memeber/config"
	"memeber/middleware/stats"
	"memeber/renderer"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// init db

	// init web framework
	e := echo.New()
	initEcho(e)
	e.Logger.Fatal(e.Start(":" + config.GetInstance().Service.Port))
}

func initEcho(e *echo.Echo) {
	// Middleware
	e.Use(stats.GetInstance().Middleware())
	e.Use(middleware.CORS())
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// Renderer
	e.Renderer = renderer.NewTemplateRanderer()

	// Rotues
	registHandler(e)
}

func registHandler(e *echo.Echo) {
	e.Static("static", "static")
	e.GET("/stats", stats.GetInstance().Handle)

	e.GET("/index", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", nil)
	})

	// TODO: swagger 연동
	//e.GET(ServiceRoot+"swagger/*", swaggerHandler.WrapHandler)
}
