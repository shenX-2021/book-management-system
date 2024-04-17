package main

import (
	"book-management-system-server/global"
	model "book-management-system-server/model"
	router "book-management-system-server/router"
	_ "book-management-system-server/validator"
)

func main() {
	global.DB = model.GetDB()
	r := router.InitRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
