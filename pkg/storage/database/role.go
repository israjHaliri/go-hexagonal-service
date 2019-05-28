package database

import "time"

type Role struct {
	ID      int       `json:"id" gorm:"primary_key"`
	Role    string    `json:"role" gorm:"unique_index:idx_role_name"`
	Created time.Time `json:"created" sql:"DEFAULT:current_timestamp"`
	Updated time.Time `json:"updated" gorm:"nullable" sql:"DEFAULT:null"`
}
