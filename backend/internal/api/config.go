package api 

import "github.com/NazarM11/TerestPin/internal/database"

type ApiConfig struct {
    DB         *database.Queries
    Uploads    string
    JwtSecret  string
}

