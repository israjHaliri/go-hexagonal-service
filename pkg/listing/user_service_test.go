package listing

import (
	"github.com/israjHaliri/go-hexagonal-service/pkg/storage/database"
	"github.com/israjHaliri/go-hexagonal-service/pkg/storage/mock"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetAllUsers(t *testing.T) {
	userRepositoryMock := new(mock.UserRepositoryMock)

	mockRoles := []database.Role{}
	mockUser := database.User{1, "israj", "israj.haliri@gmail.com", "qwerty", true, time.Now(), time.Now(), mockRoles}

	mockListUser := make([]database.User, 0)
	mockListUser = append(mockListUser, mockUser)

	userRepositoryMock.On("FindAllUser", 1).Return(mockListUser, nil).Once()

	userService := NewUserService(userRepositoryMock)
	listUser, _ := userService.GetAllUsers()

	log.Info(len(mockListUser))
	log.Info(len(listUser))

	assert.Len(t, listUser, len(mockListUser))
	userRepositoryMock.AssertExpectations(t)
}
