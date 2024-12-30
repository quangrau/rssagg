package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	// Env
	godotenv.Load(".env")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("PORT must be set")
	}

	// Router configs
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1router := chi.NewRouter()
	v1router.Get("/ready", func(rw http.ResponseWriter, req *http.Request) {
		respondWithJSON(rw, http.StatusOK, map[string]string{"status": "success"})
	})

	v1router.Get("/err", func(rw http.ResponseWriter, req *http.Request) {
		respondWithError(rw, http.StatusInternalServerError, "Internal Server Error")
	})

	router.Mount("/v1", v1router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + PORT,
	}
	log.Printf("Listening on port %s", PORT)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
