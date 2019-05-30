package listing

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/israjHaliri/go-hexagonal-service/pkg/storage/database"
)

type implement struct {
	userRepository database.UserRepository
	roleRepository database.RoleRepository
}

func NewService(userRepository database.UserRepository, roleRepository database.RoleRepository) UserService {
	return &implement{userRepository, roleRepository}
}

type UserService interface {
	GetAllUsers(page int, limit int) *pagination.Paginator
	GetUserById(id int) (User, error)
	GetAllRoles() ([]Role, error)
	GetRoleById(id int) (Role, error)
}
