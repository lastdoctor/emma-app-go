package service

type UserRepo interface {
	GetUser()
	CreateUser()
	UpdateUser()
	DeleteUser()
}
