package listing

import (
	"github.com/israjHaliri/go-hexagonal-service/pkg/storage/database"
	"github.com/israjHaliri/go-hexagonal-service/pkg/storage/mock"
	"testing"
	"time"
)

func TestGetAllUsers(t *testing.T) {
	userRepositoryMock := new(mock.UserRepositoryMock)

	mockRoles := []database.Role{}
	mockUser := database.User{1, "israj", "israj.haliri@gmail.com", "qwerty", true, time.Now(), time.Now(), mockRoles}

	mockListUser := []database.User{}
	mockListUser = append(mockListUser, mockUser)

	t.Run("success", func(t *testing.T) {
		userRepositoryMock.On("FindAllUser").Return(mockListUser, "next-cursor", nil).Once()

		userService := NewUserService(userRepositoryMock)
		userService.GetAllUsers()
	})
}
