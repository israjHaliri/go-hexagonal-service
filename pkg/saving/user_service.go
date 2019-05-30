package saving

import (
	"github.com/israjHaliri/go-hexagonal-service/pkg/storage/database"
	"time"
)

type implementUser struct {
	userRepository database.UserRepository
}

type UserService interface {
	CreateUser(user *User) (*User, error)
	UpdateUser(user *User) (*User, error)
}

func NewService(userRepository database.UserRepository) UserService {
	return &implementUser{userRepository}
}

func (userService *implementUser) CreateUser(user *User) (*User, error) {
	dbUser := database.User{}
	dbUser.Username = user.Username
	dbUser.Email = user.Email
	dbUser.Password = user.Password
	dbUser.Active = user.Active
	dbUser.Created = time.Now()

	_, err := userService.userRepository.SaveUser(dbUser)

	return user, err
}

func (userService *implementUser) UpdateUser(user *User) (*User, error) {
	panic("implement me")
}
