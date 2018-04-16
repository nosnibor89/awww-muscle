package main

import (

	"github.com/labstack/echo"

	"github.com/nosnibor89/awww-muscle/config"
)

func main() {
	e := echo.New()

	config.New(config.Framework{e})

	e.Logger.Fatal(e.Start(":3000"))
}