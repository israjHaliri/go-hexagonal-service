package database

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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

func TestDeleteUser(t *testing.T) {
	userRepository := NewUserRepository(connectionTest.GormDb)

	listUser, errListing := userRepository.FindAllUser()

	if errListing != nil {
		t.Error("Testing update user failed !", errListing)
	}

	var err []error
	for _, data := range listUser {
		err = append(err, userRepository.DeleteUser(data.ID))
	}

	if err[0] != nil || errListing != nil {
		t.Error("Testing delete user failed !", err)
	} else {
		t.Log("Testing delete user ok !")
	}

	connectionTest.GormDb.Exec("DELETE FROM user_roles")
}
