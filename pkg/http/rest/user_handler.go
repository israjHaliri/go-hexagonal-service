package rest

import (
	"github.com/israjHaliri/go-hexagonal-service/pkg/listing"
	"github.com/labstack/echo"
	"net/http"
)

type UserHandler struct {
	Lister listing.Service
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewUserHandler(e *echo.Echo, lister listing.Service) {
	handler := &UserHandler{
		Lister: lister,
	}

	e.GET("/users", handler.GetUsers)
	e.GET("/users/:id", handler.GetUsers)
	e.POST("/users", handler.GetUsers)
	e.DELETE("/users/:id", handler.GetUsers)
}

// FetchArticle will fetch the article based on given params
func (a *UserHandler) GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, a.Lister.GetAllUsers())
}
