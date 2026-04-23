package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/NazarM11/TerestPin/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	db      *database.Queries
	uploads string
	jwtSecret      string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	uploadPath := os.Getenv("UPLOAD_PATH")
	if uploadPath == "" {
		log.Fatal("UPLOAD_PATH environment variable is not set")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable is not set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable is not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	cfg := apiConfig{
		db:      dbQueries,
		uploads: uploadPath,
		jwtSecret: jwtSecret,
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Chi!"))
	})
	fs := http.FileServer(http.Dir(uploadPath))
	r.Handle("/uploads/*", http.StripPrefix("/uploads/", fs))
	
	r.Group(func(r chi.Router) {
    	r.Use(cfg.MiddlewareAuth)
    
    	r.Post("/api/pins", cfg.CreatePin)
	})
    
	r.Post("/users", cfg.CreateUser)
	r.Post("/pins", cfg.CreatePin)

	

	fmt.Println("Server is live at http://localhost:8080")
	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
