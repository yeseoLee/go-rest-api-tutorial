package main

import (
	"board/datasource"
	"board/question"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// datasource
	ds := datasource.
		MySQLInstance().
		MySQLConnectionInfo("ip", 1234, "id", "pw", "dbName").
		MySQLConnect()

	////// Echo Web Framework //////
	e := echo.New()
	// Middleware
	registMiddleware(e)
	// Routes
	registRoutes(ds, e)
	// Start server
	runServer(e)
}

func registMiddleware(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}

func registRoutes(ds datasource.DataSource, e *echo.Echo) {
	// Question
	question.RegistQuestionRoute(ds, e)
}

func runServer(e *echo.Echo) {
	e.Logger.Fatal(e.Start(":1323"))
}
