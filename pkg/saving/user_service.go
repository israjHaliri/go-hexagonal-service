package saving

import (
	"github.com/israjHaliri/go-hexagonal-service/pkg/storage/database"
	"github.com/israjHaliri/go-hexagonal-service/pkg/util"
	"time"
)

func (implement *implement) CreateUser(user *SaveUser) (*SaveUser, error) {
	dbUser := database.User{}
	dbUser.Username = user.Username
	dbUser.Email = user.Email
	pass, errHash := util.HashPassword(user.Password)
	if errHash != nil {
		return user, errHash
	}
	dbUser.Password = pass
	dbUser.Active = user.Active
	dbUser.Created = time.Now()

	_, err := implement.userRepository.SaveUser(dbUser)

	return user, err
}

func (implement *implement) UpdateUser(user *UpdateUser) (*UpdateUser, error) {
	dbUser, errGet := implement.userRepository.FindUserById(user.ID)

	if errGet != nil {
		return user, errGet
	}
	dbUser.Username = user.Username
	dbUser.Email = user.Email
	pass, errHash := util.HashPassword(user.Password)
	if errHash != nil {
		return user, errHash
	}
	dbUser.Password = pass

	dbUser.Active = user.Active
	dbUser.Updated = time.Now()

	_, err := implement.userRepository.UpdateUser(dbUser)

	return user, err
}
