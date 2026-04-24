package handlers

import (
	"net/http"
	"time"

	"github.com/NazarM11/TerestPin/internal/api"
	"github.com/NazarM11/TerestPin/internal/api/utils"
	"github.com/NazarM11/TerestPin/internal/auth"
)

type jwt_json struct {
	Token string `json:"token"`
}

func RefreshUserToken(cfg *api.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := auth.GetBearerToken(r.Header)
		if err != nil {
			utils.RespondWithError(w, 401, "Failed to read token", err)
			return
		}

		user, err := cfg.DB.GetUserByRefreshToken(r.Context(), token)
		if err != nil {
			utils.RespondWithError(w, 401, "Failed to fetch user by token", err)
			return
		}

		jwt_token, err := auth.MakeJWT(user.ID, cfg.JwtSecret, time.Hour)
		if err != nil {
			utils.RespondWithError(w, 500, "Failed to create jwt", err)
			return
		}

		utils.RespondWithJSON(w, 200, jwt_json{
			Token: jwt_token,
		})
	}
}
