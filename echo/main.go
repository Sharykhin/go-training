package main

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"go-training/echo/app/controllers"
	"net/http"
)

func hello(c *echo.Context) error {
	return c.String(http.StatusOK, "Hello World!n")
}

func main() {
	e := echo.New()

	e.Use(mw.Logger())
	e.Use(mw.Recover())

	e.Get("/", hello)
	e.Get("/users", controllers.Get)

	e.Run(":8000")
}
