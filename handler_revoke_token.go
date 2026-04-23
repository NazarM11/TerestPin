package main

import (
	"net/http"

	"github.com/NazarM11/TerestPin/internal/auth"
)

func (cfg *apiConfig) RevokeToken(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, 401, "Failed to read token", err)
		return
	}

	_, err = cfg.db.RevokeRefreshToken(r.Context(), token)
	if err != nil {
		respondWithError(w, 500, "Failed to revoke token", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
