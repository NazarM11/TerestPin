package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/NazarM11/TerestPin/internal/api"
	"github.com/NazarM11/TerestPin/internal/api/utils"
	"github.com/NazarM11/TerestPin/internal/auth"
	"github.com/NazarM11/TerestPin/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type User struct {
	ID             uuid.UUID `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"-"`
}

type parameters struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

func CreateUser(cfg *api.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		params := parameters{}
		err := decoder.Decode(&params)
		if err != nil {
			utils.RespondWithError(w, 400, "Error ocuured", err)
			return
		}

		if params.Email == "" || params.Password == "" {
			utils.RespondWithError(w, 400, "Email or password empty", nil)
			return
		}

		hashedPassword, err := auth.HashPassword(params.Password)
		if err != nil {
			utils.RespondWithError(w, 500, "Password hashing failed", err)
			return
		}

		now := time.Now().UTC()

		user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
			ID:             uuid.New(),
			CreatedAt:      sql.NullTime{Time: now, Valid: true},
			UpdatedAt:      sql.NullTime{Time: now, Valid: true},
			Email:          params.Email,
			HashedPassword: hashedPassword,
		})
		if err != nil {
			if pgErr, ok := err.(*pq.Error); ok {
				if pgErr.Code == "23505" && pgErr.Constraint == "users_email_key" {
					utils.RespondWithError(w, 409, "Email already exists", nil)
					return
				}
			}
			utils.RespondWithError(w, 500, "User creation failed", err)
			return
		}

		utils.RespondWithJSON(w, 201, databaseUserToUser(user))
	}
}
