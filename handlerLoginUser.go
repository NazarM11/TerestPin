package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/NazarM11/TerestPin/internal/auth"
	"github.com/NazarM11/TerestPin/internal/database"
)

type response struct {
	User `json:"user"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func (cfg *apiConfig) Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, "Error ocuured", err)
		return
	}

	if params.Email == "" || params.Password == "" {
		respondWithError(w, 400, "Email or password empty", nil)
		return
	}

	user, err := cfg.db.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, 401, "Invalid email or password", err)
		return
	}

	ok, err := auth.CheckPasswordHash(params.Password, user.HashedPassword)
	if err != nil || !ok {
		respondWithError(w, 401, "Invalid email or password", err)
		return
	}
	
	accesss_token, err := auth.MakeJWT(user.ID, cfg.jwtSecret, time.Hour)
	if err != nil {
		respondWithError(w, 500, "Jwt token error", err)
		return
	}

	refresh_token := auth.MakeRefreshToken()
	refreshToken, err := cfg.db.CreateRefreshToken(r.Context(), database.CreateRefreshTokenParams{
		Token:     refresh_token,
		UserID:    user.ID,
		ExpiresAt: time.Now().UTC().Add(time.Hour * 24 * 60),
	})
	if err != nil {
		respondWithError(w, 500, "Refresh error", err)
		return
	}

	cleanedUser := databaseUserToUser(user)

	resp := response{
        User:         cleanedUser,
        Token:        accesss_token,
        RefreshToken: refreshToken.Token,
    }

	respondWithJSON(w, 200, resp)
}
