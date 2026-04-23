package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/NazarM11/TerestPin/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (cfg *apiConfig) DeletePin(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(UserIDKey).(uuid.UUID)
	if !ok {
		respondWithError(w, 401, "Unauthorized", nil)
		return
	}

	pinIDString := chi.URLParam(r, "PinID")
	pinID, err := uuid.Parse(pinIDString)
	if err != nil {
		respondWithError(w, 500, "Failed to parse pin ID", err)
		return
	}

	pin, err := cfg.db.GetPin(r.Context(), pinID)
	if err != nil {
		respondWithError(w, 404, "Failed to fetch pin", err)
		return
	}

	if userID != pin.UserID {
		respondWithError(w, 403, "Current user id doesnt matсh pin creator id", nil)
		return
	}

	err = cfg.db.DeletePin(r.Context(), database.DeletePinParams{
		ID:     pin.ID,
		UserID: pin.UserID,
	})
	if err != nil {
		respondWithError(w, 500, "Failed to delete pin", err)
		return
	}

	fileName := filepath.Base(pin.ImageUrl)
	filePath := filepath.Join(cfg.uploads, fileName)

	os.Remove(filePath)
	w.WriteHeader(http.StatusNoContent)
}
