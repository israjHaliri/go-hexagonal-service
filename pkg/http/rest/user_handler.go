package rest

import (
	"github.com/israjHaliri/go-hexagonal-service/pkg/listing"
	"github.com/labstack/echo"
	"net/http"
)

type UserHandler struct {
	Lister listing.Service
}

func NewUserHandler(e *echo.Echo, lister listing.Service) {
	handler := &UserHandler{
		Lister: lister,
	}

	e.GET("/users", handler.GetUsers)
	e.GET("/users/:id", handler.GetUsers)
	e.POST("/users", handler.GetUsers)
	e.PUT("/users", handler.GetUsers)
	e.DELETE("/users/:id", handler.GetUsers)
}

func (userhandler *UserHandler) GetUsers(c echo.Context) error {
	listUser, err := userhandler.Lister.GetAllUsers()

	if err != nil {
		return c.JSON(http.StatusOK, err)
	}

	return c.JSON(http.StatusNotFound, listUser)
}
