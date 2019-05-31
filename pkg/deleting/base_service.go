package deleting

import (
	"github.com/israjHaliri/go-hexagonal-service/pkg/storage/database"
)

type implement struct {
	userRepository database.UserRepository
	roleRepository database.RoleRepository
}

func NewService(userRepository database.UserRepository, roleRepository database.RoleRepository) Service {
	return &implement{userRepository, roleRepository}
}

type Service interface {
	RemoveUser(id int) error
	RemoveRole(id int) error
}
