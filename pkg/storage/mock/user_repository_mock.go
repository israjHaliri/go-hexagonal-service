package mock

import (
	"github.com/israjHaliri/go-hexagonal-service/pkg/storage/database"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (mock *UserRepositoryMock) SaveUser(user database.User) (database.User, error) {
	//retrival := mock.Called(user)
	//
	//var return0 database.User
	//var return1 error
	//if ref, ok := retrival.Get(0).(func(*database.User) error); ok {
	//	return0 = ref(&user)
	//	return1 = ref(&user)
	//} else {
	//	return0 = retrival.Error(0)
	//}
	//
	//return return0

	panic("implement me")
}

func (mock *UserRepositoryMock) FindAllUser(total int) ([]database.User, error) {
	retrival := mock.Called(total)

	var return0 []database.User

	if ref, ok := retrival.Get(0).(func() []database.User); ok {
		return0 = ref()
	} else {
		if retrival.Get(0) != nil {
			return0 = retrival.Get(0).([]database.User)
		}
	}

	var return2 error
	if ref, ok := retrival.Get(1).(func() error); ok {
		return2 = ref()
	} else {
		return2 = retrival.Error(1)
	}

	log.Info("len users : ", len(return0))
	log.Info("user : ", return0[0])

	return retrival.Get(0).([]database.User), return2
}

func (mock *UserRepositoryMock) FindUserById(id int) (database.User, error) {
	panic("implement me")
}

func (mock *UserRepositoryMock) UpdateUser(user database.User) (database.User, error) {
	panic("implement me")
}

func (mock *UserRepositoryMock) DeleteUser(id int) error {
	panic("implement me")
}
