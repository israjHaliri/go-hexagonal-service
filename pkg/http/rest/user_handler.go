package rest

import (
	"github.com/israjHaliri/go-hexagonal-service/pkg/deleting"
	"github.com/israjHaliri/go-hexagonal-service/pkg/listing"
	"github.com/israjHaliri/go-hexagonal-service/pkg/saving"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
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
	e.GET("/users/:id", handler.GetUserById)
	e.PUT("/users", handler.UpdateUser)
	e.PUT("/users/:id/roles/:id_role", handler.UpdateUserRole)
	e.DELETE("/users/:id", handler.DeleteUser)
}

func (userhandler *UserHandler) CreateUsers(c echo.Context) error {
	userReq := new(saving.SaveUser)
	if err := c.Bind(userReq); err != nil || len(userReq.Role) < 1 {
		return c.JSON(http.StatusBadRequest, response{http.StatusBadRequest, err.Error()})
	}

	user, err := userhandler.Saver.CreateUser(userReq)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{http.StatusInternalServerError, err.Error()})
	}

	return c.JSON(http.StatusCreated, response{http.StatusCreated, user})
}

func (userhandler *UserHandler) GetUsers(c echo.Context) error {
	listUser := userhandler.Lister.GetAllUsers(1, 10)

	return c.JSON(http.StatusOK, listUser)
}

func (userhandler *UserHandler) GetUserById(c echo.Context) error {
	id := c.Param("id")

	currId, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response{http.StatusBadRequest, err.Error()})
	}

	listUser, err := userhandler.Lister.GetUserById(currId)

	if err != nil {
		return c.JSON(http.StatusNotFound, response{http.StatusNotFound, err.Error()})
	}

	return c.JSON(http.StatusOK, response{http.StatusOK, listUser})
}

func (userhandler *UserHandler) UpdateUser(c echo.Context) error {
	userReq := new(saving.UpdateUser)
	if err := c.Bind(userReq); err != nil {
		return c.JSON(http.StatusBadRequest, response{http.StatusBadRequest, err.Error()})
	}

	user, err := userhandler.Saver.UpdateUser(userReq)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{http.StatusInternalServerError, err.Error()})
	}

	return c.JSON(http.StatusOK, response{http.StatusOK, user})
}

func (userhandler *UserHandler) UpdateUserRole(c echo.Context) error {
	id := c.Param("id")
	idRole := c.Param("id_role")

	currId, err := strconv.Atoi(id)
	currRoleId, errCvrtRoleId := strconv.Atoi(idRole)

	if err != nil || errCvrtRoleId != nil {
		return c.JSON(http.StatusBadRequest, response{http.StatusBadRequest, err.Error()})
	}

	userRoleReq := new(saving.UpdateUserRole)
	if err := c.Bind(userRoleReq); err != nil {
		return c.JSON(http.StatusBadRequest, response{http.StatusBadRequest, err.Error()})
	}

	errSave := userhandler.Saver.UpdateUserRole(currId, currRoleId, userRoleReq)

	if errSave != nil {
		return c.JSON(http.StatusInternalServerError, response{http.StatusInternalServerError, err.Error()})
	}

	return c.JSON(http.StatusOK, response{http.StatusOK, "Success"})
}

func (userhandler *UserHandler) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	currId, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response{http.StatusBadRequest, err.Error()})
	}

	errDelete := userhandler.Deleter.RemoveUser(currId)

	if errDelete != nil {
		return c.JSON(http.StatusInternalServerError, response{http.StatusInternalServerError, err.Error()})
	}

	return c.JSON(http.StatusOK, response{http.StatusOK, "Deleted"})
}
