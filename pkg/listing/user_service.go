package listing

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/israjHaliri/go-hexagonal-service/pkg/storage/database"
)

type userService struct {
	userRepository database.UserRepository
}

type Service interface {
	GetAllUsers(page int, limit int) *pagination.Paginator
}

func NewUserService(userRepository database.UserRepository) Service {
	return &userService{userRepository}
}

func (userService *userService) GetAllUsers(page int, limit int) *pagination.Paginator {
	return userService.userRepository.FindAllUser(page, limit)
}
