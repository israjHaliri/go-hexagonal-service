package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Connection struct {
	GormDb *gorm.DB
}

type Repository interface {
	FindAllUser() ([]User, error)
	FindById(id int) (User, error)
	Save(user User) (User, error)
	Update(user User) (User, error)
	Delete(id int) error
}

func NewUserRepository(gormDB *gorm.DB) Repository {
	return &Connection{gormDB}
}

func (conn *Connection) FindAllUser() ([]User, error) {
	listUser := []User{}

	db := conn.GormDb
	defer db.Close()

	err := db.Find(&listUser).Error

	return listUser, err
}

func (conn *Connection) FindById(id int) (User, error) {
	db := conn.GormDb
	defer db.Close()

	user := User{}

	err := db.Where("id = ?", id).First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (conn *Connection) Save(user User) (User, error) {
	db := conn.GormDb
	defer db.Close()

	err := db.Create(&user).Error

	return user, err
}

func (conn *Connection) Update(user User) (User, error) {
	db := conn.GormDb
	defer db.Close()

	err := db.Save(&user).Error

	return user, err
}

func (conn *Connection) Delete(id int) error {
	db := conn.GormDb
	defer db.Close()

	err := db.Delete(&User{}, "id = ?", id).Error

	return err
}
