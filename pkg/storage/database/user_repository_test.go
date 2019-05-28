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
	connectionDatabase := config.NewMysqlConnectionDatabase()
	gormDB := connectionDatabase.Open()
	gormDB.AutoMigrate(&User{}, &Role{})

	if err := gormDB.Table("user_roles").AddForeignKey("role_id", "roles (id)", "CASCADE", "CASCADE").Error; err != nil {
		gormDB.DropTable(&User{}, &Role{}, "user_roles")

		panic(err)
	}

	connectionTest.GormDb = gormDB

	code := m.Run()

	connectionTest.GormDb = nil

	gormDB.Close()

	os.Exit(code)
}

func TestSaveUser(t *testing.T) {
	userRepository := NewUserRepository(connectionTest.GormDb)

	user := User{}
	user.Username = "israjj"
	user.Password = "12345678"
	user.Email = "israj.haliri@gmail.com"
	user.Active = true
	user.Created = time.Now()

	role := Role{}
	role.Role = "AMDIN"

	listRoles := []Role{}
	listRoles = append(listRoles, role)

	user.Roles = listRoles

	_, err := userRepository.SaveUser(user)

	if err != nil {
		t.Error("Testing save user failed !", err)
	} else {
		t.Log("Testing save user ok !")
	}
}

func TestFindAllUser(t *testing.T) {
	userRepository := NewUserRepository(connectionTest.GormDb)

	_, err := userRepository.FindAllUser()

	if err != nil {
		t.Log(" Testing find all users failed !", err)
	} else {
		t.Log("Testing find all users ok !")
	}
}

func TestFindUserById(t *testing.T) {
	userRepository := NewUserRepository(connectionTest.GormDb)

	listUser, errListing := userRepository.FindAllUser()

	if errListing != nil {
		t.Error("Testing update user failed !", errListing)
	}

	_, err := userRepository.FindUserById(listUser[0].ID)

	if err != nil {
		t.Error("Testing find by id user failed !", err)
	} else {
		t.Log("Testing find by id user ok !")
	}
}

func TestUpdateUser(t *testing.T) {
	userRepository := NewUserRepository(connectionTest.GormDb)

	listUser, errListing := userRepository.FindAllUser()

	if errListing != nil {
		t.Error("Testing update user failed !", errListing)
	}

	user := listUser[0]
	user.Username = "israj h"
	user.Password = "12345678"
	user.Email = "israj.haliri@gmail.com"
	user.Active = true
	user.Updated = time.Now()

	_, err := userRepository.UpdateUser(user)

	if err != nil {
		t.Error("Testing update user failed !", err)
	} else {
		t.Log("Testing update user ok !")
	}
}

//func TestDeleteUser(t *testing.T) {
//	userRepository := NewUserRepository(connectionTest.GormDb)
//
//	listUser, errListing := userRepository.FindAllUser()
//
//	if errListing != nil {
//		t.Error("Testing update user failed !", errListing)
//	}
//
//	err := userRepository.DeleteUser(listUser[0].ID)
//
//	if err != nil || errListing != nil {
//		t.Error("Testing delete user failed !")
//	} else {
//		t.Log("Testing delete user ok !")
//	}
//}
