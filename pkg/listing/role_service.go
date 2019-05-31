package listing

func (implement *implement) GetAllRoles() ([]Role, error) {
	currentListRole, err := implement.roleRepository.FindAllRole()

	listRole := []Role{}
	for _, data := range currentListRole {
		role := Role{}
		role.ID = data.ID
		role.Role = data.Role
		role.Created = data.Created
		role.Updated = data.Updated

		listRole = append(listRole, role)
	}

	return listRole, err
}

func (implement *implement) GetRoleById(id int) (*Role, error) {
	currentRole, err := implement.roleRepository.FindRoleById(id)

	role := Role{}
	role.ID = currentRole.ID
	role.Role = currentRole.Role
	role.Created = currentRole.Created
	role.Updated = currentRole.Updated

	return &role, err
}
