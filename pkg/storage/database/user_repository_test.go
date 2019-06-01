package database

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"reflect"
	"testing"
	"time"
)

func TestSaveUser(t *testing.T) {
	userRepository := NewUserRepository(connectionTest.GormDb)

	user := User{}
	user.Username = "israj"
	user.Password = "12345678"
	user.Email = "israj.haliri@gmail.com"
	user.Active = true
	user.Created = time.Now()

	roleRepository := NewRoleRepository(connectionTest.GormDb)

	listRoles, _ := roleRepository.FindAllRole()

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

	paginator := userRepository.FindAllUser(1, 10)

	if paginator.TotalRecord < 1 {
		t.Log(" Testing find all users failed !")
	} else {
		t.Log("Testing find all users ok ! ")
	}
}

func TestFindUserById(t *testing.T) {
	userRepository := NewUserRepository(connectionTest.GormDb)

	paginator := userRepository.FindAllUser(1, 10)

	list := reflect.ValueOf(paginator.Records).Interface().(*[]User)

	user := (*list)[0]

	_, err := userRepository.FindUserById(user.ID)

	if err != nil {
		t.Error("Testing find by id user failed !", err)
	} else {
		t.Log("Testing find by id user ok !")
	}
}

func TestUpdateUser(t *testing.T) {
	userRepository := NewUserRepository(connectionTest.GormDb)

	paginator := userRepository.FindAllUser(1, 10)

	list := reflect.ValueOf(paginator.Records).Interface().(*[]User)

	user := (*list)[0]

	user, err := userRepository.FindUserById(user.ID)

	user.Username = "israj h"
	user.Password = "12345678"
	user.Email = "israj.haliri@gmail.com"
	user.Active = true
	timeLocal := time.Now()
	user.Updated = &timeLocal

	_, errUpdate := userRepository.UpdateUser(user)

	if errUpdate != nil {
		t.Error("Testing update user failed !", err)
	} else {
		t.Log("Testing update user ok !")
	}
}

func TestDeleteUser(t *testing.T) {
	userRepository := NewUserRepository(connectionTest.GormDb)

	paginator := userRepository.FindAllUser(1, 10)

	list := reflect.ValueOf(paginator.Records).Interface().(*[]User)

	var err []error
	for _, data := range *list {
		err = append(err, userRepository.DeleteUser(data.ID))
	}

	if err[0] != nil {
		t.Error("Testing delete user failed !", err)
	} else {
		t.Log("Testing delete user ok !")
	}

	connectionTest.GormDb.Exec("DELETE FROM user_roles")
}
