package rest

import (
	"github.com/israjHaliri/go-hexagonal-service/pkg/listing"
	"github.com/israjHaliri/go-hexagonal-service/pkg/saving"
	"github.com/labstack/echo"
	"net/http"
)

type UserHandler struct {
	Lister listing.UserService
	Saver  saving.UserService
}

func NewUserHandler(e *echo.Echo, lister listing.UserService, saver saving.UserService) {
	handler := &UserHandler{
		Lister: lister,
		Saver:  saver,
	}

	e.GET("/users", handler.GetUsers)
	e.GET("/users/:id", handler.GetUsers)
	e.POST("/users", handler.CreateUsers)
	e.PUT("/users", handler.GetUsers)
	e.DELETE("/users/:id", handler.GetUsers)
}

func (userhandler *UserHandler) GetUsers(c echo.Context) error {
	listUser := userhandler.Lister.GetAllUsers(1, 10)

	return c.JSON(http.StatusOK, listUser)
}

func (userhandler *UserHandler) CreateUsers(c echo.Context) error {
	user, err := userhandler.Saver.CreateUser(&saving.User{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, user)
}
