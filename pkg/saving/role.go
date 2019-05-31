package saving

type SaveRole struct {
	Role string `json:"role"`
}

type SaveRoleByID struct {
	ID int `json:"id"`
}

type UpdateRole struct {
	ID   int    `json:"id"`
	Role string `json:"role"`
}
