package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Connection struct {
	GormDb *gorm.DB
}

type Repository interface {
	FindAllUser() []User
}

func NewUserRepository(gormDB *gorm.DB) Repository {
	return &Connection{gormDB}
}

func (conn *Connection) FindAllUser() []User {
	listUser := []User{}

	db := conn.GormDb
	defer db.Close()

	db.Find(&listUser)

	return listUser
}
