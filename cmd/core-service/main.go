package main

import (
	"github.com/israjHaliri/go-hexagonal-service/pkg/config"
	"github.com/israjHaliri/go-hexagonal-service/pkg/deleting"
	"github.com/israjHaliri/go-hexagonal-service/pkg/http/rest"
	"github.com/israjHaliri/go-hexagonal-service/pkg/listing"
	"github.com/israjHaliri/go-hexagonal-service/pkg/saving"
	"github.com/israjHaliri/go-hexagonal-service/pkg/storage/database"
	"github.com/labstack/echo"
)

func main() {
	connectionDatabase := config.NewMysqlConnectionDatabase()
	gormDB := connectionDatabase.Open()

	defer gormDB.Close()

	e := echo.New()

	userRepository := database.NewUserRepository(gormDB)
	roleRepository := database.NewRoleRepository(gormDB)

	listingService := listing.NewService(userRepository, roleRepository)
	savingService := saving.NewService(userRepository, roleRepository)
	deleteService := deleting.NewService(userRepository, roleRepository)

	rest.NewUserHandler(e, listingService, savingService, deleteService)
	rest.NewRoleHandler(e, listingService, savingService, deleteService)

	e.Logger.Fatal(e.Start(":10000"))
}
