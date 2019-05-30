package saving

import (
	"github.com/israjHaliri/go-hexagonal-service/pkg/storage/database"
	"github.com/israjHaliri/go-hexagonal-service/pkg/storage/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestSaveUser(t *testing.T) {
	userRepositoryMock := new(mocks.UserRepositoryMock)

	var mockUser database.User
	mockUser.ID = 1
	mockUser.Username = "israj"
	mockUser.Email = "israj.haliri@gmail.com"
	mockUser.Password = "12345678"
	mockUser.Active = true
	mockUser.Created = time.Now()
	mockUser.Updated = time.Now()

	userRepositoryMock.On("SaveUser", mock.AnythingOfType("User")).Return(mockUser, nil).Once()

	userService := NewUserService(userRepositoryMock)

	var user User
	user.Username = "israj"
	user.Email = "israj.haliri@gmail.com"
	user.Password = "12345678"
	user.Active = true

	result, err := userService.CreateUser(&user)

	assert.Equal(t, result.Email, mockUser.Email)
	assert.Equal(t, err, nil)
	userRepositoryMock.AssertExpectations(t)
}
