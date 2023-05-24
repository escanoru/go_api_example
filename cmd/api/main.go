package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

const (
	version = "1.0.0"
)

type config struct {
	port        int
	environment string
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.environment, "environment", "Dev", "Environment (dev|stage|prod)")
	flag.Parse()
	mux := http.NewServeMux() // As a best practice create your own ServeMux instance instead of using the default ServeMux which is activated by passing nil on http.ListenAndServe(port, nil), this is because the devault ServerMux uses global varables which can be used to inject handlers, by creating our own instance we prevent this.
	mux.HandleFunc("/v1/healthcheck", healthcheck)
	log.Printf("Starting server on all interfaces at port %s", p)
	log.Fatal(http.ListenAndServe(port, mux))
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, fmt.Sprintf("Error: %s %s", r.Method, http.StatusText(http.StatusMethodNotAllowed)), http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", environment)
	fmt.Fprintf(w, "version: %s\n", version)
}
