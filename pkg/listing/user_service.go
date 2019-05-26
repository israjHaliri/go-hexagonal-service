package listing

import (
	"github.com/israjHaliri/go-hexagonal-service/pkg/storage/database"
)

type userService struct {
	userRepository database.Repository
}

type Service interface {
	GetAllUsers() []User
}

func NewUserService(userRepository database.Repository) Service {
	return &userService{userRepository}
}

func (userService *userService) GetAllUsers() []User {
	listUserDatabase := userService.userRepository.FindAllUser()

	users := []User{}

	for _, data := range listUserDatabase {
		user := User{}
		user.Name = data.Name
		users = append(users, user)
	}

	return users
}
