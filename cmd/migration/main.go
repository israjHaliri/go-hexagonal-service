package main

import (
	"fmt"
	"github.com/israjHaliri/go-hexagonal-service/pkg/config"
	"github.com/israjHaliri/go-hexagonal-service/pkg/storage/database"
)

func main() {
	connectionDatabase := config.NewMysqlConnectionDatabase()
	gormDB := connectionDatabase.Open()

	defer gormDB.Close()

	gormDB.AutoMigrate(&database.User{}, &database.Role{})

	if err := gormDB.
		Table("user_roles").AddForeignKey("role_id", "roles (id)", "CASCADE", "CASCADE").
		Table("user_roles").AddForeignKey("user_id", "users (id)", "CASCADE", "CASCADE").Error; err != nil {

		gormDB.DropTable(&database.User{}, &database.Role{}, "user_roles")

		panic(err)
	}

	fmt.Println("Migration success !")
}
