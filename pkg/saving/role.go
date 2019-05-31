package saving

type SaveRole struct {
	Role string `json:"role" gorm:"unique_index:idx_role_name"`
}

type UpdateRole struct {
	ID   int    `json:"id"`
	Role string `json:"role" gorm:"unique_index:idx_role_name"`
}
