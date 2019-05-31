package saving

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
	CreateUser(user *SaveUser) (*SaveUser, error)
	UpdateUser(user *UpdateUser) (*UpdateUser, error)
	UpdateUserRole(userId int, roleId int, userRole *UpdateUserRole) error
	CreateRole(role *SaveRole) (*SaveRole, error)
	UpdateRole(role *UpdateRole) (*UpdateRole, error)
}
