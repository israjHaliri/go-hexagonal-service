package listing

import (
	"github.com/israjHaliri/go-hexagonal-service/pkg/storage/database"
	"time"
)

type userService struct {
	userRepository database.Repository
	contextTimeout time.Duration
}

func NewUserService(userRepository database.Repository, contextTimeout time.Duration) Service {
	return &userService{userRepository, contextTimeout}
}

type Service interface {
	GetAllUsers() []User
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
