package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	port, environment, version = ":4000", "DEV", "1.0.0"
)

func main() {
	mux := http.NewServeMux() // As a best practice create your own ServeMux instance instead of using the default ServeMux which is activated by passing nil on http.ListenAndServe(port, nil), this is because the devault ServerMux uses global varables which can be used to inject handlers, by creating our own instance we prevent this.
	mux.HandleFunc("/v1/healthcheck", healthcheck)
	log.Printf("Starting server on all interfaces at port %s", port)
	log.Fatal(http.ListenAndServe(port, mux))
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", environment)
	fmt.Fprintf(w, "version: %s\n", version)
}
