package main

import (
	"github.com/israjHaliri/go-hexagonal-service/pkg/http/rest"
	"github.com/israjHaliri/go-hexagonal-service/pkg/listing"
	"github.com/israjHaliri/go-hexagonal-service/pkg/storage/database"
	"github.com/labstack/echo"
	"time"
)

func main() {
	e := echo.New()

	userRepository := database.NewUserRepository("")

	timeoutContext := time.Duration(1) * time.Second

	userService := listing.NewUserService(userRepository, timeoutContext)

	rest.NewUserHandler(e, userService)

	e.Logger.Fatal(e.Start(":10000"))
}
