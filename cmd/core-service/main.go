package main

import (
	"github.com/labstack/echo"
	"github.com/israjHaliri/go-hexagonal-service/pkg/http/rest"
)

func main() {
	e := echo.New()
	e.GET("/users", rest.GetUsers)
	e.Logger.Fatal(e.Start(":10000"))
}