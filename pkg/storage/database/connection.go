package database

import "github.com/jinzhu/gorm"

type Connection struct {
	GormDb *gorm.DB
}

var connectionTest struct {
	GormDb *gorm.DB
}
