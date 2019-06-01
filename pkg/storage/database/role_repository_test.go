package database

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"testing"
	"time"
)

func TestSaveRole(t *testing.T) {
	roleRepository := NewRoleRepository(connectionTest.GormDb)

	role := Role{}
	role.Role = "ADMIN"
	role.Created = time.Now()

	_, err := roleRepository.SaveRole(role)

	roleGuest := Role{}
	roleGuest.Role = "GUEST"
	roleGuest.Created = time.Now()

	_, errRoleGuest := roleRepository.SaveRole(roleGuest)

	if err != nil || errRoleGuest != nil {
		t.Error("Testing save role failed !", err)
	} else {
		t.Log("Testing save role ok !")
	}
}

func TestFindAllRole(t *testing.T) {
	roleRepository := NewRoleRepository(connectionTest.GormDb)

	_, err := roleRepository.FindAllRole()

	if err != nil {
		t.Log(" Testing find all roles failed !", err)
	} else {
		t.Log("Testing find all roles ok !")
	}
}

func TestFindRoleById(t *testing.T) {
	roleRepository := NewRoleRepository(connectionTest.GormDb)

	listRole, errListing := roleRepository.FindAllRole()

	if errListing != nil {
		t.Error("Testing update role failed !", errListing)
	}

	_, err := roleRepository.FindRoleById(listRole[0].ID)

	if err != nil {
		t.Error("Testing find by id role failed !", err)
	} else {
		t.Log("Testing find by id role ok !")
	}
}

func TestUpdateRole(t *testing.T) {
	roleRepository := NewRoleRepository(connectionTest.GormDb)

	listRole, errListing := roleRepository.FindAllRole()

	if errListing != nil {
		t.Error("Testing update role failed !", errListing)
	}

	role := listRole[0]
	role.Role = "SUPER_ADMIN"
	timeLocal := time.Now()
	role.Updated = &timeLocal

	_, err := roleRepository.UpdateRole(role)

	if err != nil {
		t.Error("Testing update role failed !", err)
	} else {
		t.Log("Testing update role ok !")
	}
}

func TestDeleteRole(t *testing.T) {
	roleRepository := NewRoleRepository(connectionTest.GormDb)

	listRole, errListing := roleRepository.FindAllRole()

	if errListing != nil {
		t.Error("Testing update role failed !", errListing)
	}

	var err []error
	for _, data := range listRole {
		err = append(err, roleRepository.DeleteRole(data.ID))
	}

	if err[0] != nil || errListing != nil {
		t.Error("Testing delete role failed !")
	} else {
		t.Log("Testing delete role ok !")
	}
}
