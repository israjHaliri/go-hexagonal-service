package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Mysql struct {
}

type SqLite struct {
}

type ConnectionDatabase interface {
	Open() *gorm.DB
}

func NewMysqlConnectionDatabase() ConnectionDatabase {
	return &Mysql{}
}

func NewSqliteConnectionDatabase() ConnectionDatabase {
	return &SqLite{}
}

func (mysql Mysql) Open() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/go-hexagonal-service?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	return db
}

func (sqlite SqLite) Open() *gorm.DB {
	db, err := gorm.Open("sqlite3", "/tmp/go-hexagonal-service.db")

	if err != nil {
		panic(err)
	}

	return db
}
