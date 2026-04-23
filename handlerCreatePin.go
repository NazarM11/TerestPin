package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/NazarM11/TerestPin/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) CreatePin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(50 << 20)
	if err != nil {
		fmt.Printf("Multipart Error: %v\n", err)
		respondWithError(w, 400, "File too large", err)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		respondWithError(w, 400, "Invalid image", err)
		return
	}
	defer file.Close()

	title := r.FormValue("title")
	userIDStr := r.FormValue("user_id") // We'll get real auth later

	fileExt := filepath.Ext(header.Filename)
	newFilename := fmt.Sprintf("%s%s", uuid.New().String(), fileExt)
	fullPath := filepath.Join(cfg.uploads, newFilename)

	dst, err := os.Create(fullPath)
	if err != nil {
		respondWithError(w, 500, "Couldn't save file", err)
		return
	}
	defer dst.Close()
	io.Copy(dst, file)

	userID, _ := uuid.Parse(userIDStr)
	now := time.Now().UTC()

	pin, err := cfg.db.CreatePin(r.Context(), database.CreatePinParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Title:     title,
		ImageUrl:  fullPath,
		UserID:    userID,
	})

	if err != nil {
		respondWithError(w, 500, "Database error", err)
		return
	}

	respondWithJSON(w, 201, pin)
}
