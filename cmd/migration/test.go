package main

import (
	"fmt"
	"github.com/israjHaliri/go-hexagonal-service/pkg/config"
	"github.com/israjHaliri/go-hexagonal-service/pkg/storage/database"
)

func main() {
	connectionDatabase := config.NewSqliteConnectionDatabase()
	gormDB := connectionDatabase.Open()

	defer gormDB.Close()

	gormDB.AutoMigrate(&database.User{}, &database.Role{})

	fmt.Println("Migration success !")
}
