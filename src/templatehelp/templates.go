package templatehelp

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

// BuildPage build page for template.
func BuildPage(w http.ResponseWriter, t string, i interface{}) {
	ti, err := template.New("").Funcs(template.FuncMap{
		"secend": func(i, j int) bool { return i%j == j-1 },
	}).ParseFiles(
		t,
		"templates/base.html",
	)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	bta := strings.Split(t, "/")
	bt := bta[len(bta)-1]
	err = ti.ExecuteTemplate(w, bt, i)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
