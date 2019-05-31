package listing

import (
	"time"
)

type User struct {
	ID       int        `json:"id"`
	Username string     `json:"username"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
	Active   bool       `json:"active"`
	Created  time.Time  `json:"created"`
	Updated  *time.Time `json:"updated"`
}
