package saving

import (
	"github.com/israjHaliri/go-hexagonal-service/pkg/storage/database"
	"time"
)

func (implement *implement) CreateRole(role *SaveRole) (*SaveRole, error) {
	dbRole := database.Role{}
	dbRole.Role = role.Role
	dbRole.Created = time.Now()

	_, err := implement.roleRepository.SaveRole(dbRole)

	return role, err
}

func (implement *implement) UpdateRole(role *UpdateRole) (*UpdateRole, error) {
	dbRole, errGet := implement.roleRepository.FindRoleById(role.ID)

	if errGet != nil {
		return role, errGet
	}
	dbRole.Role = role.Role
	t := time.Now()
	dbRole.Updated = &t

	_, err := implement.roleRepository.UpdateRole(dbRole)

	return role, err
}
