package main

import (
	"net/http"
	"github.com/NazarM11/TerestPin/internal/auth"
	"time"
)

type jwt_json struct {
	Token string `json:"token"`
}

func (cfg *apiConfig) RefreshUserToken(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, 401, "Failed to read token", err)
		return
	}

	user, err := cfg.db.GetUserByRefreshToken(r.Context(), token)
	if err != nil {
		respondWithError(w, 401, "Failed to fetch user by token", err)
		return 
	}

	jwt_token, err := auth.MakeJWT(user.ID, cfg.jwtSecret, time.Hour)
	if err != nil {
		respondWithError(w, 500, "Failed to create jwt", err)
		return 
	}

	respondWithJSON(w, 200, jwt_json{
		Token: jwt_token,
	})
}