package handlers

import (
	"net/http"

	"github.com/NazarM11/TerestPin/internal/api"
	"github.com/NazarM11/TerestPin/internal/api/utils"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func GetPinsByUserID(cfg *api.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userIDString := chi.URLParam(r, "UserID")
		userID, err := uuid.Parse(userIDString)
		if err != nil {
			utils.RespondWithError(w, 500, "Failed to parse user ID", err)
			return
		}

		dbPins, err := cfg.DB.GetPinsByUserID(r.Context(), userID)
		if err != nil {
			utils.RespondWithError(w, 500, "Failed to fetch this user's pins", err)
			return
		}

		pins := []PinResponse{}
		for _, dbPin := range dbPins {
			pins = append(pins, PinResponse{
				ID:        dbPin.ID,
				CreatedAt: dbPin.CreatedAt,
				UpdatedAt: dbPin.UpdatedAt,
				Title:     dbPin.Title,
				ImageUrl:  dbPin.ImageUrl,
				UserID:    dbPin.UserID,
			})
		}

		utils.RespondWithJSON(w, 200, pins)
	}
}
