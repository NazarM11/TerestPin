package handlers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/NazarM11/TerestPin/internal/api"
	"github.com/NazarM11/TerestPin/internal/api/middleware"
	"github.com/NazarM11/TerestPin/internal/api/utils"
	"github.com/NazarM11/TerestPin/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func DeletePin(cfg *api.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value(middleware.UserIDKey).(uuid.UUID)
		if !ok {
			utils.RespondWithError(w, 401, "Unauthorized", nil)
			return
		}

		pinIDString := chi.URLParam(r, "PinID")
		pinID, err := uuid.Parse(pinIDString)
		if err != nil {
			utils.RespondWithError(w, 500, "Failed to parse pin ID", err)
			return
		}

		pin, err := cfg.DB.GetPin(r.Context(), pinID)
		if err != nil {
			utils.RespondWithError(w, 404, "Failed to fetch pin", err)
			return
		}

		if userID != pin.UserID {
			utils.RespondWithError(w, 403, "Current user id doesnt matсh pin creator id", nil)
			return
		}

		err = cfg.DB.DeletePin(r.Context(), database.DeletePinParams{
			ID:     pin.ID,
			UserID: pin.UserID,
		})
		if err != nil {
			utils.RespondWithError(w, 500, "Failed to delete pin", err)
			return
		}

		fileName := filepath.Base(pin.ImageUrl)
		filePath := filepath.Join(cfg.Uploads, fileName)

		os.Remove(filePath)
		w.WriteHeader(http.StatusNoContent)
	}
}
