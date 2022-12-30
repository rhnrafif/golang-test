package controllers

import (
	"go-payment/model"
	utils "go-payment/utils"
	"log"
	"net/http"

	sqlserver "go-payment/database"

	"github.com/gorilla/mux"
)

func getCustomerFromDb() []model.Customer {

	var Customer []model.Customer
	sqlserver.Database.Find(&Customer)
	return Customer

}

func allCustomer(w http.ResponseWriter, c *http.Request) {

	utils.JsonResponse(w, getCustomerFromDb())
}

func getCustomerById(w http.ResponseWriter, c *http.Request) {

}

func HandleRequest() {

	myrouter := mux.NewRouter().StrictSlash(false)

	myrouter.HandleFunc("/customer", allCustomer).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", myrouter))
}
