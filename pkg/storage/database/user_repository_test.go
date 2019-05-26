package database

import (
	"github.com/israjHaliri/go-hexagonal-service/pkg/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
	"testing"
	"time"
)

var connectionTest struct {
	GormDb *gorm.DB
}

func TestMain(m *testing.M) {
	connectionDatabase := config.NewSqliteConnectionDatabase()
	gormDB := connectionDatabase.Open()
	gormDB.AutoMigrate(User{})

	connectionTest.GormDb = gormDB

	code := m.Run()

	connectionTest.GormDb = nil

	gormDB.Close()

	os.Exit(code)
}

func TestSave(t *testing.T) {
	userRepository := NewUserRepository(connectionTest.GormDb)

	user := User{}
	user.Name = "israj"
	user.Created = time.Now()

	_, err := userRepository.Save(user)

	if err != nil {
		t.Error("Testing save user failed !")
	} else {
		t.Log("Testing save user ok !")
	}
}

func TestUpdate(t *testing.T) {
	userRepository := NewUserRepository(connectionTest.GormDb)

	listUsers, _ := userRepository.FindAllUser()

	user := listUsers[0]
	user.Name = "Jono Haliri"

	_, err := userRepository.Update(user)

	if err != nil {
		t.Error("Testing update user failed !")
	} else {
		t.Log("Testing update user ok !")
	}
}

//func TestFindAll(t *testing.T) {
//	gormDB := setupConnection()
//
//	userRepository := NewUserRepository(gormDB)
//
//	listUsers, err := userRepository.FindAllUser()
//
//	if len(listUsers) >= 1 && err != nil {
//		t.Log("Find all users ok !")
//	}
//}
//
//func TestFindById(t *testing.T) {
//	gormDB := setupConnection()
//
//	userRepository := NewUserRepository(gormDB)
//
//	listUsers, err := userRepository.FindAllUser()
//
//	if len(listUsers) >= 1 && err != nil {
//		t.Log("Find all users ok !")
//	}
//}
//
//func TestDelete(t *testing.T) {
//	gormDB := setupConnection()
//
//	userRepository := NewUserRepository(gormDB)
//
//	listUsers, err := userRepository.FindAllUser()
//
//	if len(listUsers) >= 1 && err != nil {
//		t.Log("Find2 all users ok !")
//	}
//}
