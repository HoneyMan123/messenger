package main

import (
	"net/http"

	"main.go/commonFunctions"
)

func main() {
	router := commonFunctions.Routing()
	http.ListenAndServe(":8080", router)

}
