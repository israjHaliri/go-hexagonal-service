package database

import "time"

type User struct {
	ID       int       `gorm:"primary_key" json:"id"`
	Username string    `json:"username" gorm:"unique_index:idx_username_code""`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Active   bool      `json:"active"`
	Created  time.Time `json:"created" sql:"DEFAULT:current_timestamp"`
	Updated  time.Time `json:"updated" sql:"DEFAULT:current_timestamp"`
}
