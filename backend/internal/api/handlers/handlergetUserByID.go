package handlers

import (
	"net/http"
	"time"

	"github.com/NazarM11/TerestPin/internal/api"
	"github.com/NazarM11/TerestPin/internal/api/middleware"
	"github.com/NazarM11/TerestPin/internal/api/utils"
	"github.com/google/uuid"
)

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

func GetUserByID(cfg *api.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value(middleware.UserIDKey).(uuid.UUID)
		if !ok {
			utils.RespondWithError(w, 401, "Unauthorized", nil)
			return
		}

		user, err := cfg.DB.GetUser(r.Context(), userID)
		if err != nil {
			utils.RespondWithError(w, 500, "Unable to get user", err)
			return
		}

		utils.RespondWithJSON(w, 200, UserResponse{
			ID:        user.ID,
			CreatedAt: user.CreatedAt.Time,
			UpdatedAt: user.UpdatedAt.Time,
			Email:     user.Email,
		})
	}
}
