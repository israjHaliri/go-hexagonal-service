package saving

type SaveRole struct {
	Role string `json:"role"`
}

type UpdateRole struct {
	ID   int    `json:"id"`
	Role string `json:"role"`
}
