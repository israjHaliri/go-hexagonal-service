package database

import "time"

type User struct {
	ID       int        `json:"id" gorm:"primary_key"`
	Username string     `json:"username" gorm:"unique_index:idx_username_code"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
	Active   bool       `json:"active"`
	Created  time.Time  `json:"created"`
	Updated  *time.Time `json:"updated" gorm:"nullable" sql:"DEFAULT:null"`
	Roles    []Role     `gorm:"many2many:user_roles;"`
}
