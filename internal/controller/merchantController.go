package controller

import (
	"encoding/json"
	"github.com/lastdoctor/emma-app-go/internal/logger"
	"go.uber.org/zap"
	"log"
	"net/http"
)

type MerchantCon interface {
	GetMerchantById(http.ResponseWriter, *http.Request)
	CreateMerchant(http.ResponseWriter, *http.Request)
	UpdateMerchant(http.ResponseWriter, *http.Request)
	DeleteMerchant(http.ResponseWriter, *http.Request)
}

//type MerchantController struct {
//	service service.MerchantService
//}
type CreateM struct {
	DisplayName string `json:"displayName"`
	IconUrl     string `json:"iconUrl"`
	FunnyGifUrl string `json:"funnyGifUrl"`
}

// Manual testing routes works or not.
func CreateMerchant(w http.ResponseWriter, r *http.Request) {
	var Merchant CreateM
	err := json.NewDecoder(r.Body).Decode(&Merchant)
	if err != nil {
		logger.Logger().Error("Failed to decode JSON", zap.Error(err))
	}
	log.Println(Merchant.DisplayName)
	WriteJSON(w, 200, &Merchant, nil)
}

func GetMerchantById(w http.ResponseWriter, r *http.Request) {}
func UpdateMerchant(w http.ResponseWriter, r *http.Request)  {}
func DeleteMerchant(w http.ResponseWriter, r *http.Request)  {}
