package main

import (
	"tma.com.vn/lngochuy/rest/models"
	"tma.com.vn/lngochuy/rest/routers"
)

func main() {
	models.InitializeDB()
	models.CreateTable(true)

	router := routers.Route()
	router.Run(":1999")
}
