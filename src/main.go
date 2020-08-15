package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/tomcuzz/Termovision-Frontend/src/pages"
)

var (
	server_address string
	backend_addresses string
)

func main() {
	// Parse flags
	parseFlags()

	// Build pages
	fmt.Println("Setting up pages")
	buildhandlers()

	// Run server
	fmt.Println("Stating server on: ", server_address)
	log.Fatal(http.ListenAndServe(server_address, nil))
}

func parseFlags() {
	flag.StringVar(&server_address, "server_address", ":8080", "Address to run webserver on")
	flag.StringVar(&backend_addresses, "backend_addresses", "", "Comma seperated list of backend server addresses (can have host:port if needed)")

	flag.Parse()
}

func buildhandlers() {
	http.HandleFunc("/", pages.HomeHandler)
}
