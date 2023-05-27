package main

import (
	"log"
	"net/http"
	"volume-assignment/pkg/server"
)

func main() {
	mux := http.NewServeMux()

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	srv := server.New()
	mux.HandleFunc("/calculate", srv.CalculateResponse)
	log.Fatal(s.ListenAndServe())
}
