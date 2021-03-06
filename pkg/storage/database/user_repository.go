package database

import (
	"github.com/israjHaliri/go-hexagonal-service/pkg/util"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserRepository interface {
	SaveUser(user User) (User, error)
	FindAllUser(page int, limit int) *util.Paginator
	FindUserById(id int) (User, error)
	FindUserByContext(coloumn string, value string) (User, error)
	UpdateUser(user User) (User, error)
	UpdateUserRole(userId int, roleIdExisting int, roleIdNew int) error
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

func (conn *Connection) FindAllUser(page int, limit int) *util.Paginator {
	var listUser []User

	db := conn.GormDb

	paginator := util.Paging(&util.ParamPaging{
		DB:      db,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"id desc"},
		ShowSQL: false,
	}, &listUser, util.TypeUser)

	return paginator
}

func (conn *Connection) FindUserByContext(coloumn string, value string) (User, error) {
	db := conn.GormDb

	user := User{}

	err := db.Preload("Roles").First(&user, "username = ?", value).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (conn *Connection) FindUserById(id int) (User, error) {
	db := conn.GormDb

	user := User{}

	err := db.Preload("Roles").First(&user, "id = ?", id).Error

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

func (conn *Connection) UpdateUserRole(userId int, roleIdExisting int, roleIdNew int) error {
	db := conn.GormDb

	err := db.Table("user_roles").Where("user_id = ? AND role_id = ?", userId, roleIdExisting).Update("role_id", roleIdNew).Error

	return err
}

func (conn *Connection) DeleteUser(id int) error {
	db := conn.GormDb

	err := db.Delete(&User{}, "id = ?", id).Error

	return err
}
