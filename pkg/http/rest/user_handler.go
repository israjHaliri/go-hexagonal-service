package rest

import (
	"github.com/israjHaliri/go-hexagonal-service/pkg/deleting"
	"github.com/israjHaliri/go-hexagonal-service/pkg/listing"
	"github.com/israjHaliri/go-hexagonal-service/pkg/saving"
	"github.com/labstack/echo"
	"net/http"
)

type UserHandler struct {
	Lister  listing.Service
	Saver   saving.Service
	Deleter deleting.Service
}

func NewUserHandler(e *echo.Echo, lister listing.Service, saver saving.Service, deleter deleting.Service) {
	handler := &UserHandler{
		Lister:  lister,
		Saver:   saver,
		Deleter: deleter,
	}

	e.POST("/users", handler.CreateUsers)
	e.GET("/users", handler.GetUsers)
	e.GET("/users/:id", handler.GetUsers)
	e.PUT("/users", handler.GetUsers)
	e.DELETE("/users/:id", handler.GetUsers)
}

func (userhandler *UserHandler) CreateUsers(c echo.Context) error {
	userReq := new(saving.SaveUser)
	if err := c.Bind(userReq); err != nil || len(userReq.Role) < 1 {
		return c.JSON(http.StatusBadRequest, err)
	}

	user, err := userhandler.Saver.CreateUser(&saving.SaveUser{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, user)
}

func (userhandler *UserHandler) GetUsers(c echo.Context) error {
	listUser := userhandler.Lister.GetAllUsers(1, 10)

	return c.JSON(http.StatusOK, listUser)
}
