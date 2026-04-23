package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func (cfg *apiConfig) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(UserIDKey).(uuid.UUID)
	if !ok {
		respondWithError(w, 401, "Unauthorized", nil)
		return
	}

	pins, err := cfg.db.GetPinsByUserID(r.Context(), userID)
	if err != nil {
		respondWithError(w, 500, "Failed to get pins for this user", err)
		return
	}

	for _, pin := range pins {
		fileName := filepath.Base(pin.ImageUrl)
		filePath := filepath.Join(cfg.uploads, fileName)

		os.Remove(filePath)
	}

	err = cfg.db.DeleteUser(r.Context(), userID)
	if err != nil {
		respondWithError(w, 500, "Failed to delete user", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
