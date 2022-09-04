package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.Int("port", 8080, "port to listen on")
	flag.Parse()

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	log.Printf("listening on port %d", *port)

	err := http.ListenAndServe(
		fmt.Sprintf("127.0.0.1:%d", *port),
		mux,
	)
	if err != nil {
		log.Fatalf("Error listening for traffic: %v", err)
	}
}
