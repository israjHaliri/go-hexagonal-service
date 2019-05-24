package database

import "time"

type Connection struct {
	Url string
}

type Repository interface {
	FindAllUser() []User
}

func NewUserRepository(Url string) Repository {
	return &Connection{Url}
}

func (conn *Connection) FindAllUser() []User {
	listUser := []User{}
	listUser = append(listUser, User{1, "Israj", time.Now()})

	return listUser
}
