package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type RoleRepository interface {
	SaveRole(role Role) (Role, error)
	FindAllRole() ([]Role, error)
	FindRoleById(id int) (Role, error)
	UpdateRole(role Role) (Role, error)
	DeleteRole(id int) error
}

func NewRoleRepository(gormDB *gorm.DB) RoleRepository {
	return &Connection{gormDB}
}

func (conn *Connection) SaveRole(role Role) (Role, error) {
	db := conn.GormDb

	err := db.Create(&role).Error

	return role, err
}

func (conn *Connection) FindAllRole() ([]Role, error) {
	listRole := []Role{}

	db := conn.GormDb

	err := db.Find(&listRole).Error

	return listRole, err
}

func (conn *Connection) FindRoleById(id int) (Role, error) {
	db := conn.GormDb

	role := Role{}

	err := db.Where("id = ?", id).First(&role).Error

	if err != nil {
		return role, err
	}

	return role, nil
}

func (conn *Connection) UpdateRole(role Role) (Role, error) {
	db := conn.GormDb

	err := db.Save(&role).Error

	return role, err
}

func (conn *Connection) DeleteRole(id int) error {
	db := conn.GormDb

	err := db.Delete(&Role{}, "id = ?", id).Error

	return err
}
