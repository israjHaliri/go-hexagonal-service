package listing

import (
	"github.com/israjHaliri/go-hexagonal-service/pkg/storage/database"
)

type userService struct {
	userRepository database.Repository
}

type Service interface {
	GetAllUsers() ([]User, error)
}

func NewUserService(userRepository database.Repository) Service {
	return &userService{userRepository}
}

func (userService *userService) GetAllUsers() ([]User, error) {
	listUserDatabase, err := userService.userRepository.FindAllUser()

	listUser := []User{}

	if len(listUserDatabase) > 0 && err != nil {
		for _, data := range listUserDatabase {
			user := User{}
			user.Name = data.Name
			listUser = append(listUser, user)
		}
	}

	return listUser, nil
}
