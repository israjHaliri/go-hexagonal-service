package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserRepository interface {
	SaveUser(user User) (User, error)
	FindAllUser(total int) ([]User, error)
	FindUserById(id int) (User, error)
	UpdateUser(user User) (User, error)
	DeleteUser(id int) error
}

func NewUserRepository(gormDB *gorm.DB) UserRepository {
	return &Connection{gormDB}
}

func (conn *Connection) SaveUser(user User) (User, error) {
	db := conn.GormDb

	err := db.Create(&user).Error

	return user, err
}

func (conn *Connection) FindAllUser(total int) ([]User, error) {
	listUser := []User{}

	db := conn.GormDb

	err := db.Find(&listUser).Error

	return listUser, err
}

func (conn *Connection) FindUserById(id int) (User, error) {
	db := conn.GormDb

	user := User{}

	err := db.Where("id = ?", id).First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (conn *Connection) UpdateUser(user User) (User, error) {
	db := conn.GormDb

	err := db.Save(&user).Error

	return user, err
}

func (conn *Connection) DeleteUser(id int) error {
	db := conn.GormDb

	err := db.Delete(&User{}, "id = ?", id).Error

	return err
}
