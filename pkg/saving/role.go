package saving

type Role struct {
	Role string `json:"role" gorm:"unique_index:idx_role_name"`
}
