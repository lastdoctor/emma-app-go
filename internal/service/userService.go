package service

import "fmt"

type UserService interface {
	GetUser()
	CreateUser()
	UpdateUser()
	DeleteUser()
}

func CreateUser(data interface{}) {
	fmt.Print(data)
}
