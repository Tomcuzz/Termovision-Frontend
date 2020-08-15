package pages

import (
	"net/http"

	"github.com/tomcuzz/Termovision-Frontend/src/templatehelp"
)

// HomeHandler is a web handler for the home page.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	a := []Section{}
	t := "templates/home.html"
	templatehelp.BuildPage(w, t, a)
}

type Item struct {
	Text string
}

type Section struct {
	Title string
	Items []Item
}
