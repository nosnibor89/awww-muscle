package config

import "github.com/labstack/echo"

type Framework struct {
	*echo.Echo
}

func New(framework Framework)  {
	SetRoutes(framework)
}
