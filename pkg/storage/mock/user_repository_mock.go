package mock

import (
	"github.com/israjHaliri/go-hexagonal-service/pkg/storage/database"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stretchr/testify/mock"
)

type userRepositoryMock struct {
	mock.Mock
}

func (mock *userRepositoryMock) SaveUser(user database.User) (database.User, error) {
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

func (mock *userRepositoryMock) FindAllUser() ([]database.User, error) {
	panic("implement me")
}

func (mock *userRepositoryMock) FindUserById(id int) (database.User, error) {
	panic("implement me")
}

func (mock *userRepositoryMock) UpdateUser(user database.User) (database.User, error) {
	panic("implement me")
}

func (mock *userRepositoryMock) DeleteUser(id int) error {
	panic("implement me")
}
