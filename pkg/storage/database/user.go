package database

import "time"

type User struct {
	ID      int       `gorm:"primary_key" json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
}
