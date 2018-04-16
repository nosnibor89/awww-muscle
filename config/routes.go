package config

import (
	"github.com/labstack/echo"
	"net/http"
)

func SetRoutes(f Framework)  {
	f.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	f.GET("/test/:name", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, " + c.Param("name"))
	})

}