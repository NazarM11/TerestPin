package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/NazarM11/TerestPin/internal/api"
	"github.com/NazarM11/TerestPin/internal/api/utils"
	"github.com/NazarM11/TerestPin/internal/auth"
	"github.com/NazarM11/TerestPin/internal/database"
)

type response struct {
	User         `json:"user"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func LoginUser(cfg *api.ApiConfig) http.HandlerFunc {
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

		user, err := cfg.DB.GetUserByEmail(r.Context(), params.Email)
		if err != nil {
			utils.RespondWithError(w, 401, "Invalid email or password", err)
			return
		}

		ok, err := auth.CheckPasswordHash(params.Password, user.HashedPassword)
		if err != nil || !ok {
			utils.RespondWithError(w, 401, "Invalid email or password", err)
			return
		}

		accesss_token, err := auth.MakeJWT(user.ID, cfg.JwtSecret, time.Hour)
		if err != nil {
			utils.RespondWithError(w, 500, "Jwt token error", err)
			return
		}

		refresh_token := auth.MakeRefreshToken()
		refreshToken, err := cfg.DB.CreateRefreshToken(r.Context(), database.CreateRefreshTokenParams{
			Token:     refresh_token,
			UserID:    user.ID,
			ExpiresAt: time.Now().UTC().Add(time.Hour * 24 * 60),
		})
		if err != nil {
			utils.RespondWithError(w, 500, "Refresh error", err)
			return
		}

		cleanedUser := databaseUserToUser(user)

		resp := response{
			User:         cleanedUser,
			Token:        accesss_token,
			RefreshToken: refreshToken.Token,
		}

		utils.RespondWithJSON(w, 200, resp)
	}
}
