package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"strings"

	"github.com/NazarM11/TerestPin/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) CreatePin(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(UserIDKey).(uuid.UUID)
	if !ok {
		respondWithError(w, 401, "Unauthorized", nil)
		return
	}

	err := r.ParseMultipartForm(50 << 20)
	if err != nil {
		respondWithError(w, 400, "File too large", err)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		respondWithError(w, 400, "Invalid image", err)
		return
	}
	defer file.Close()

	fileExt := filepath.Ext(header.Filename)
	allowedExts := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".gif": true}
	
	if !allowedExts[strings.ToLower(fileExt)] {
    	respondWithError(w, 400, "Invalid file type", nil)
    	return
	}

	newFilename := fmt.Sprintf("%s%s", uuid.New().String(), fileExt)
	fullPath := filepath.Join(cfg.uploads, newFilename)

	dst, err := os.Create(fullPath)
	if err != nil {
		respondWithError(w, 500, "Couldn't save file", err)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		respondWithError(w, 500, "Failed to save file contents", err)
		return
	}

	publicURL := fmt.Sprintf("/uploads/%s", newFilename)
	now := time.Now().UTC()

	pin, err := cfg.db.CreatePin(r.Context(), database.CreatePinParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Title:     r.FormValue("title"),
		ImageUrl:  publicURL,
		UserID:    userID,
	})

	if err != nil {
		respondWithError(w, 500, "Database error", err)
		return
	}

	respondWithJSON(w, 201, pin)
}
