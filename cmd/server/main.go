package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./assets/")))

	log.Println(os.LookupEnv("PORT"))

	srv := &http.Server{
		Addr:    os.Getenv("PORT"),
		Handler: mux,
	}

	log.Printf("Serving\n")

	log.Fatal(srv.ListenAndServe())

}
