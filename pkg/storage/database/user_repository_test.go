package database

import (
	"github.com/israjHaliri/go-hexagonal-service/pkg/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"testing"
)

func setup() *gorm.DB {
	connectionDatabase := config.NewMysqlConnectionDatabase()
	gormDB := connectionDatabase.Open()
	gormDB.AutoMigrate(User{})

	return gormDB
}

func TestFetch(t *testing.T) {
	gormDB := setup()

	userRepository := NewUserRepository(gormDB)

	listUsers, err := userRepository.FindAllUser()

	if len(listUsers) >= 1 && err != nil {
		t.Log("Find all users ok !")
	}
}
