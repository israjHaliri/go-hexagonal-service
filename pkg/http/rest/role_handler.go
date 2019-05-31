package rest

import (
	"github.com/israjHaliri/go-hexagonal-service/pkg/deleting"
	"github.com/israjHaliri/go-hexagonal-service/pkg/listing"
	"github.com/israjHaliri/go-hexagonal-service/pkg/saving"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type RoleHandler struct {
	Lister  listing.Service
	Saver   saving.Service
	Deleter deleting.Service
}

func NewRoleHandler(e *echo.Echo, lister listing.Service, saver saving.Service, deleter deleting.Service) {
	handler := &RoleHandler{
		Lister:  lister,
		Saver:   saver,
		Deleter: deleter,
	}

	e.POST("/roles", handler.CreateRoles)
	e.GET("/roles", handler.GetRoles)
	e.GET("/roles/:id", handler.GetRolesById)
	e.PUT("/roles", handler.UpdateRole)
	e.DELETE("/roles/:id", handler.DeleteRole)
}

func (rolehandler *RoleHandler) CreateRoles(c echo.Context) error {
	roleReq := new(saving.SaveRole)
	if err := c.Bind(roleReq); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	role, err := rolehandler.Saver.CreateRole(roleReq)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{http.StatusInternalServerError, err.Error()})
	}

	return c.JSON(http.StatusCreated, response{http.StatusCreated, role})
}

func (rolehandler *RoleHandler) GetRoles(c echo.Context) error {
	listRole, err := rolehandler.Lister.GetAllRoles()

	if err != nil {
		return c.JSON(http.StatusNotFound, response{http.StatusNotFound, listRole})
	}

	return c.JSON(http.StatusOK, response{http.StatusOK, listRole})
}

func (rolehandler *RoleHandler) GetRolesById(c echo.Context) error {
	id := c.Param("id")

	currId, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response{http.StatusBadRequest, err.Error()})
	}

	listRole, err := rolehandler.Lister.GetRoleById(currId)

	if err != nil {
		return c.JSON(http.StatusNotFound, response{http.StatusNotFound, err.Error()})
	}

	return c.JSON(http.StatusOK, response{http.StatusOK, listRole})
}

func (rolehandler *RoleHandler) UpdateRole(c echo.Context) error {
	roleReq := new(saving.UpdateRole)
	if err := c.Bind(roleReq); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	role, err := rolehandler.Saver.UpdateRole(roleReq)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{http.StatusInternalServerError, err.Error()})
	}

	return c.JSON(http.StatusOK, response{http.StatusOK, role})
}

func (rolehandler *RoleHandler) DeleteRole(c echo.Context) error {
	id := c.Param("id")

	currId, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response{http.StatusBadRequest, err.Error()})
	}

	errDelete := rolehandler.Deleter.RemoveRole(currId)

	if errDelete != nil {
		return c.JSON(http.StatusInternalServerError, response{http.StatusInternalServerError, err.Error()})
	}

	return c.JSON(http.StatusOK, response{http.StatusOK, "Deleted"})
}