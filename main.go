package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	port := "8080"

	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	log.Printf("Starting up on http://localhost:%s", port)

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", getListOfBook)
	r.Post("/posts", schedulePickup)

	log.Fatal(http.ListenAndServe(":"+port, r))
}

func getListOfBook(w http.ResponseWriter, r *http.Request) {
}

func schedulePickup(w http.ResponseWriter, r *http.Request) {
}
