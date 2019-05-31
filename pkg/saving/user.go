package saving

type SaveUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Active   bool   `json:"active"`
	Role     []SaveRoleByID
}

type UpdateUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Active   bool   `json:"active"`
	Role     []SaveRoleByID
}

type UpdateUserRole struct {
	ID int `json:"new_role_id"`
}
