package commonFunctions

import (
	"net/http"
	"text/template"
)

func ReturnTemplate(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("pages/main page/front end/html/index.html"))
	tmpl.Execute(w, nil)
}
