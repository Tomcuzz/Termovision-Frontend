package templatehelp

import (
	"html/template"
	"log"
	"net/http"
)

// BuildPage build page for template.
func BuildPage(w http.ResponseWriter, t string, i interface{}) {
	ti, err := template.ParseFiles(
		t,
		"templates/base.html",
	)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ti.Execute(w, i)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
