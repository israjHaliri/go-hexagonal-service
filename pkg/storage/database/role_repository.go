package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type RoleRepository interface {
	SaveRole(user Role) (Role, error)
	FindAllRole() ([]Role, error)
	FindRoleById(id int) (Role, error)
	UpdateRole(user Role) (Role, error)
	DeleteRole(id int) error
}

func NewRoleRepository(gormDB *gorm.DB) RoleRepository {
	return &Connection{gormDB}
}

func (conn *Connection) SaveRole(user Role) (Role, error) {
	db := conn.GormDb
	defer db.Close()

	err := db.Create(&user).Error

	return user, err
}

func (conn *Connection) FindAllRole() ([]Role, error) {
	listRole := []Role{}

	db := conn.GormDb
	defer db.Close()

	err := db.Find(&listRole).Error

	return listRole, err
}

func (conn *Connection) FindRoleById(id int) (Role, error) {
	db := conn.GormDb
	defer db.Close()

	user := Role{}

	err := db.Where("id = ?", id).First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (conn *Connection) UpdateRole(user Role) (Role, error) {
	db := conn.GormDb
	defer db.Close()

	err := db.Save(&user).Error

	return user, err
}

func (conn *Connection) DeleteRole(id int) error {
	db := conn.GormDb
	defer db.Close()

	err := db.Delete(&Role{}, "id = ?", id).Error

	return err
}
