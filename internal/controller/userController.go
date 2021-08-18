package controller

import (
	"encoding/json"
	"github.com/lastdoctor/emma-app-go/internal/logger"
	"github.com/lastdoctor/emma-app-go/internal/service"
	"go.uber.org/zap"
	"log"
	"net/http"
)

type UserCon interface {
	GetUserById(http.ResponseWriter, *http.Request)
	CreateUser(http.ResponseWriter, *http.Request)
	UpdateUser(http.ResponseWriter, *http.Request)
	DeleteUser(http.ResponseWriter, *http.Request)
}
type UserController struct {
	service service.UserService
}
type CreateU struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user CreateU
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		logger.Logger().Error("Failed to decode JSON", zap.Error(err))
	}
	log.Println(user.LastName)
	service.CreateUser(user)
	WriteJSON(w, 200, &user, nil)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {}
func UpdateUser(w http.ResponseWriter, r *http.Request)  {}
func DeleteUser(w http.ResponseWriter, r *http.Request)  {}
