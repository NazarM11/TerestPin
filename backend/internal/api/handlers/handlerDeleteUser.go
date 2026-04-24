package handlers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/NazarM11/TerestPin/internal/api"
	"github.com/NazarM11/TerestPin/internal/api/middleware"
	"github.com/NazarM11/TerestPin/internal/api/utils"
	"github.com/google/uuid"
)

func DeleteUser(cfg *api.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value(middleware.UserIDKey).(uuid.UUID)
		if !ok {
			utils.RespondWithError(w, 401, "Unauthorized", nil)
			return
		}

		pins, err := cfg.DB.GetPinsByUserID(r.Context(), userID)
		if err != nil {
			utils.RespondWithError(w, 500, "Failed to get pins for this user", err)
			return
		}

		for _, pin := range pins {
			fileName := filepath.Base(pin.ImageUrl)
			filePath := filepath.Join(cfg.Uploads, fileName)

			os.Remove(filePath)
		}

		err = cfg.DB.DeleteUser(r.Context(), userID)
		if err != nil {
			utils.RespondWithError(w, 500, "Failed to delete user", err)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
