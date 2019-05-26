package main

import (
	"github.com/israjHaliri/go-hexagonal-service/pkg/config"
	"github.com/israjHaliri/go-hexagonal-service/pkg/http/rest"
	"github.com/israjHaliri/go-hexagonal-service/pkg/listing"
	"github.com/israjHaliri/go-hexagonal-service/pkg/storage/database"
	"github.com/labstack/echo"
)

func main() {
	connectionDatabase := config.NewMysqlConnectionDatabase()
	gormDB := connectionDatabase.Open()
	gormDB.AutoMigrate(database.User{})

	e := echo.New()

	userRepository := database.NewUserRepository(gormDB)

	userService := listing.NewUserService(userRepository)

	rest.NewUserHandler(e, userService)

	e.Logger.Fatal(e.Start(":10000"))
}
