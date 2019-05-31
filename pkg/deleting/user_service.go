package deleting

func (implement *implement) RemoveUser(id int) error {
	return implement.userRepository.DeleteUser(id)
}
