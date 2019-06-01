package listing

import (
	"github.com/israjHaliri/go-hexagonal-service/pkg/storage/mocks"
	"github.com/israjHaliri/go-hexagonal-service/pkg/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestGetAllUsers(t *testing.T) {
	userRepositoryMock := new(mocks.UserRepositoryMock)

	var mockPaginator = new(util.Paginator)

	mockPaginator.TotalRecord = 10
	mockPaginator.Records = []User{}
	mockPaginator.Page = 1

	mockPaginator.Offset = 0
	mockPaginator.Limit = 10
	mockPaginator.TotalPage = 1

	userRepositoryMock.On("FindAllUser", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(mockPaginator).Once()

	service := NewService(userRepositoryMock, nil)
	resultPaginator := service.GetAllUsers(1, 10)

	assert.Equal(t, resultPaginator.TotalRecord, mockPaginator.TotalRecord)
	userRepositoryMock.AssertExpectations(t)
}
