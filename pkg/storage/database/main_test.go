package database

import (
	"github.com/israjHaliri/go-hexagonal-service/pkg/config"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	connectionDatabase := config.NewSqliteConnectionDatabase()
	gormDB := connectionDatabase.Open()

	defer gormDB.Close()

	gormDB.AutoMigrate(&User{}, &Role{})

	connectionTest.GormDb = gormDB

	code := m.Run()

	connectionTest.GormDb = nil

	gormDB.Close()

	os.Exit(code)
}
