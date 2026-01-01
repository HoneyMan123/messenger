package main

import (
	"main.go/commonFunctions"
	"net/http"
)

//func handler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "go fuck yourself antivirus")

//}

func main() {
	http.HandleFunc("/", commonFunctions.ReturnTemplate)
	http.ListenAndServe(":8080", nil)
}
