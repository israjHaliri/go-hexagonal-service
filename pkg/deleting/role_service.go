package deleting

func (implement *implement) RemoveRole(id int) error {
	return implement.roleRepository.DeleteRole(id)
}
