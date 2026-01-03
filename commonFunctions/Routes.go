package commonFunctions

import (
	"github.com/gorilla/mux"
)

func Routing() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", ReturnTemplate("pages/main page/front end/html/index.html"))
	return r
}
