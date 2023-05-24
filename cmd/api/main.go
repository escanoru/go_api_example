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
	http.HandleFunc("/v1/healthcheck", healthcheck)
	log.Printf("Starting server on all interfaces at port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", environment)
	fmt.Fprintf(w, "version: %s\n", version)
}
