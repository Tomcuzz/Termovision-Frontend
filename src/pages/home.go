package pages

import (
	"net/http"

	"github.com/tomcuzz/vennpiere/src/templatehelp"
)

// HomeHandler is a web handler for the home page.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	a = []string{}
	t := "templates/home.html"
	templatehelp.BuildPage(w, t, a)
}
