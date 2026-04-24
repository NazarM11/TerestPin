package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/NazarM11/TerestPin/internal/api"
    "github.com/NazarM11/TerestPin/internal/api/handlers"
    mymiddleware "github.com/NazarM11/TerestPin/internal/api/middleware" 
    "github.com/NazarM11/TerestPin/internal/database"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware" 
    "github.com/joho/godotenv"
    _ "github.com/lib/pq"
)
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

	err = db.Ping()
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}
	dbQueries := database.New(db)

	cfg := api.ApiConfig{
		DB:        dbQueries,
		Uploads:   uploadPath,
		JwtSecret: jwtSecret,
	}

	r := chi.NewRouter()
	r.Use(mymiddleware.MiddlewareLogger)
	r.Use(middleware.Recoverer)

	r.Post("/api/users", handlers.CreateUser(&cfg))
	r.Post("/api/login", handlers.LoginUser(&cfg))
	r.Post("/api/refresh", handlers.RefreshUserToken(&cfg))
	r.Get("/api/pins", handlers.GetPins(&cfg))
	r.Get("/api/users/{UserID}/pins", handlers.GetPinsByUserID(&cfg))

	r.Group(func(r chi.Router) {
		r.Use(mymiddleware.MiddlewareAuth(&cfg))

		r.Post("/api/pins", handlers.CreatePin(&cfg))
		r.Delete("/api/pins/{PinID}", handlers.DeletePin(&cfg))
		r.Delete("/api/users", handlers.DeleteUser(&cfg))
	})

	fs := http.FileServer(http.Dir(uploadPath))
	r.Handle("/uploads/*", http.StripPrefix("/uploads/", fs))

	fmt.Printf("Server is live at http://localhost:%s\n", port)
	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
