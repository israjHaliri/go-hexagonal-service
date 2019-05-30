package listing

import (
	"github.com/biezhi/gorm-paginator/pagination"
)

func (implement *implement) GetAllUsers(page int, limit int) *pagination.Paginator {
	return implement.userRepository.FindAllUser(page, limit)
}

func (implement *implement) GetUserById(id int) (User, error) {
	currentUser, err := implement.userRepository.FindUserById(id)

	user := User{}
	user.ID = currentUser.ID
	user.Username = currentUser.Username
	user.Email = currentUser.Email
	user.Password = currentUser.Password
	user.Active = currentUser.Active
	user.Created = currentUser.Created
	user.Updated = currentUser.Updated

	return user, err
}
