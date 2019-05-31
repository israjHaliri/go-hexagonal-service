package rest

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/israjHaliri/go-hexagonal-service/pkg/deleting"
	"github.com/israjHaliri/go-hexagonal-service/pkg/listing"
	"github.com/israjHaliri/go-hexagonal-service/pkg/saving"
	"github.com/israjHaliri/go-hexagonal-service/pkg/util"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"strconv"
	"time"
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

	e.POST("/login", handler.Login)

	r := e.Group("/users")
	r.Use(middleware.JWT([]byte("MYSECRETTOCHANG3")))

	r.POST("", handler.CreateUsers)
	r.GET("", handler.GetUsers, isAdmin)
	r.GET("/:id", handler.GetUserById)
	r.PUT("", handler.UpdateUser)
	r.PUT("/:id/roles/:id_role", handler.UpdateUserRole)
	r.DELETE("/:id", handler.DeleteUser)
}

func isAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		isAdmin := claims["admin"].(bool)
		role := claims["role"]
		fmt.Println("ROLE : ", role)
		if isAdmin == false {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}

func (userhandler *UserHandler) Login(c echo.Context) error {
	userReq := new(listing.User)
	if err := c.Bind(userReq); err != nil {
		return c.JSON(http.StatusBadRequest, response{http.StatusBadRequest, err.Error()})
	}

	user, errCheck := userhandler.Lister.GetUserByContext("username", userReq.Username)

	if errCheck != nil {
		return c.JSON(http.StatusOK, response{http.StatusOK, "Username not found"})
	} else if util.CheckPasswordHash(userReq.Password, user.Password) == false {
		return c.JSON(http.StatusUnauthorized, response{http.StatusUnauthorized, "Unauthorized"})
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.Username
	claims["admin"] = true
	claims["role"] = "guna"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("MYSECRETTOCHANG3"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response{http.StatusOK, t})
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

	return c.JSON(http.StatusOK, response{http.StatusOK, listUser})
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
