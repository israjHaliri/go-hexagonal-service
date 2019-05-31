package listing

import "time"

type Role struct {
	ID      int        `json:"id"`
	Role    string     `json:"role"`
	Created time.Time  `json:"created"`
	Updated *time.Time `json:"updated"`
}
