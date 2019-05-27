package listing

import (
	"github.com/israjHaliri/go-hexagonal-service/pkg/storage/database"
)

type userService struct {
	userRepository database.UserRepository
}

type Service interface {
	GetAllUsers() ([]User, error)
}

func NewUserService(userRepository database.UserRepository) Service {
	return &userService{userRepository}
}

func (userService *userService) GetAllUsers() ([]User, error) {
	listUserDatabase, err := userService.userRepository.FindAllUser()

	listUser := []User{}

	if len(listUserDatabase) > 0 && err != nil {
		for _, data := range listUserDatabase {
			user := User{}
			user.ID = data.ID
			user.Username = data.Username
			user.Password = data.Password
			user.Email = data.Email
			user.Active = data.Active
			user.Created = data.Created
			user.Updated = data.Updated
			listUser = append(listUser, user)
		}
	}

	return listUser, nil
}
