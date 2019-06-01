package mocks

import (
	"github.com/israjHaliri/go-hexagonal-service/pkg/storage/database"
	"github.com/israjHaliri/go-hexagonal-service/pkg/util"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (mock *UserRepositoryMock) SaveUser(user database.User) (database.User, error) {
	retrival := mock.Called(user)

	var return0 database.User
	var return1 error

	if ref, ok := retrival.Get(0).(func(database.User) database.User); ok {
		return0 = ref(user)
	} else {
		if retrival.Get(0) != nil {
			return0 = retrival.Get(0).(database.User)
		}
	}

	if ref, ok := retrival.Get(1).(func(database.User) error); ok {
		return1 = ref(user)
	} else {
		return1 = retrival.Error(1)
	}

	return return0, return1
}

func (mock *UserRepositoryMock) FindAllUser(page int, limit int) *util.Paginator {
	retrival := mock.Called(page, limit)

	var return0 *util.Paginator

	if ref, ok := retrival.Get(0).(func() *util.Paginator); ok {
		return0 = ref()
	} else {
		if retrival.Get(0) != nil {
			return0 = retrival.Get(0).(*util.Paginator)
		}
	}

	return return0
}

func (mock *UserRepositoryMock) FindUserById(id int) (database.User, error) {
	panic("implement me")
}

func (mock *UserRepositoryMock) FindUserByContext(coloumn string, value string) (database.User, error) {
	panic("implement me")
}

func (mock *UserRepositoryMock) UpdateUser(user database.User) (database.User, error) {
	panic("implement me")
}

func (mock *UserRepositoryMock) UpdateUserRole(userId int, roleIdExisting int, roleIdNew int) error {
	panic("implement me")
}

func (mock *UserRepositoryMock) DeleteUser(id int) error {
	panic("implement me")
}
