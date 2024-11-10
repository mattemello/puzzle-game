package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./web")))

	srv := &http.Server{
		Addr:    ":3001",
		Handler: mux,
	}

	log.Printf("Serving on port 3000\n")

	log.Fatal(srv.ListenAndServe())

}
