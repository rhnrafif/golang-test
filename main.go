package main

import (
	"go-payment/controllers"
	sqlserver "go-payment/database"
)

func main() {

	sqlserver.Init()
	controllers.HandleRequest()
}
