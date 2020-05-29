package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tomcuzz/Termovision-Frontend/src/pages"
)

func main() {
	fmt.Println("Setting up pages")
	buildhandlers()
	fmt.Println("Stating server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func buildhandlers() {
	http.HandleFunc("/", pages.HomeHandler)
}
