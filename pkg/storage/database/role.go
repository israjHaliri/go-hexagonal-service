package database

import "time"

type Role struct {
	ID      int       `gorm:"primary_key" json:"id"`
	Role    string    `json:"role" gorm:"unique_index:idx_rolen_code""`
	Created time.Time `json:"created" sql:"DEFAULT:current_timestamp"`
	Updated time.Time `json:"updated" sql:"DEFAULT:current_timestamp"`
}
