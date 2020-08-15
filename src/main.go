package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/tomcuzz/Termovision-Frontend/src/pages"
)

var (
	backend_addresses string
)

func main() {
	// Parse flags
	parseFlags()

	// Build pages
	fmt.Println("Setting up pages")
	buildhandlers()

	// Run server
	fmt.Println("Stating server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func parseFlags() {
	flag.StringVar(&backend_addresses, "backend_addresses", "", "Comma seperated list of backend server addresses (can have host:port if needed)")

	flag.Parse()
}

func buildhandlers() {
	http.HandleFunc("/", pages.HomeHandler)
}
