package mocks

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/israjHaliri/go-hexagonal-service/pkg/storage/database"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (mock *UserRepositoryMock) SaveUser(user database.User) (database.User, error) {
	//retrival := mocks.Called(user)
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

func (mock *UserRepositoryMock) FindAllUser(page int, limit int) *pagination.Paginator {
	retrival := mock.Called(page, limit)

	var return0 *pagination.Paginator

	if ref, ok := retrival.Get(0).(func() *pagination.Paginator); ok {
		return0 = ref()
	} else {
		if retrival.Get(0) != nil {
			return0 = retrival.Get(0).(*pagination.Paginator)
		}
	}

	return return0
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
