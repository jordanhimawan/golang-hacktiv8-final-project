package main

import (
	"net/http"
	"sesi-final-project/database"
	"sesi-final-project/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()

	http.ListenAndServe(":9000", r)
}
