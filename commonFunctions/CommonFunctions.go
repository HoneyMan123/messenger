package commonFunctions

import (
	"net/http"
	"text/template"
)

func ReturnTemplate(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles(path))
		tmpl.Execute(w, nil)
	}
}
