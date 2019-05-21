package rest

import (
	"net/http"

	"github.com/israjHaliri/go-hexagonal-service/pkg/listing"
	"github.com/labstack/echo"
)

func GetUsers(c echo.Context) error {
	// User ID from path `users/:id`
	// id := c.Param("id")

	list := listing.GetUsers()

	return c.JSON(http.StatusOK, list)
}
