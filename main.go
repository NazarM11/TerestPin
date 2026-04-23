package main

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"os"
	"log"
	""
)

type apiConfig struct {
	db			*database.Queries
	uploads 	string
}

func main() {
	godotenv.Load(".env")

	uploadPath := os.Getenv("UPLOAD_PATH")
	if uploadPath == "" {
		log.Fatal("UPLOAD_PATH environment variable is not set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable is not set")
	}

	cfg := apiConfig{
		db: db,
		uploads: uploadPath,
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)    
	r.Use(middleware.Recoverer) 

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Chi!"))
	})

	r.Get("/user/{userID}", func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "userID")
		w.Write([]byte("Looking for user: " + userID))
	})

	http.ListenAndServe(port, r)
}