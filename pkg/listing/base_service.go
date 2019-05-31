package listing

import (
	"github.com/israjHaliri/go-hexagonal-service/pkg/storage/database"
	"github.com/israjHaliri/go-hexagonal-service/pkg/util"
)

type implement struct {
	userRepository database.UserRepository
	roleRepository database.RoleRepository
}

func NewService(userRepository database.UserRepository, roleRepository database.RoleRepository) Service {
	return &implement{userRepository, roleRepository}
}

type Service interface {
	GetAllUsers(page int, limit int) *util.Paginator
	GetUserById(id int) (User, error)
	GetUserByContext(coloumn string, value string) (UserRole, error)
	GetAllRoles() ([]Role, error)
	GetRoleById(id int) (*Role, error)
}
