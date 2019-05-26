package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Mysql struct {
}

type ConnectionDatabase interface {
	Open() *gorm.DB
}

func NewMysqlConnectionDatabase() ConnectionDatabase {
	return &Mysql{}
}

func (mysql Mysql) Open() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/go-hexagonal-service?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	return db
}
