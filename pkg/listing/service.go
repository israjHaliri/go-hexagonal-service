package listing

import (
	"errors"
)

// ErrNotFound is used when a beer could not be found.
var ErrNotFound = errors.New("beer not found")

type Repository interface {
	GetAllUsers() []User
}

// Service provides beer and review listing operations.
type Service interface {
	GetUsers() []User
}

type service struct {
	r Repository
}

// NewService creates a listing service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetUsers() []User {
	return s.r.GetAllUsers()
}