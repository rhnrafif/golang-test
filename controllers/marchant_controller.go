package controllers

import (
	"go-payment/model"
	utils "go-payment/utils"
	"net/http"

	sqlserver "go-payment/database"

	"github.com/gorilla/mux"
)

func getMerchantByName(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["Name"]

	var merchant model.Merchant

	result := sqlserver.Database.Where("Name = ?", name).First(&merchant)
	if result.Error == nil {
		utils.JsonResponse(w, merchant)
	}
}
