package handlers

import (
	"net/http"

	"github.com/NazarM11/TerestPin/internal/api"
	"github.com/NazarM11/TerestPin/internal/api/utils"
	"github.com/NazarM11/TerestPin/internal/auth"
)

func RevokeToken(cfg *api.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := auth.GetBearerToken(r.Header)
		if err != nil {
			utils.RespondWithError(w, 401, "Failed to read token", err)
			return
		}

		_, err = cfg.DB.RevokeRefreshToken(r.Context(), token)
		if err != nil {
			utils.RespondWithError(w, 500, "Failed to revoke token", err)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
