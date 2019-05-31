package saving

type SaveUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Active   bool   `json:"active"`
	Role     []SaveRole
}

type UpdateUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Active   bool   `json:"active"`
	Role     []SaveRole
}
