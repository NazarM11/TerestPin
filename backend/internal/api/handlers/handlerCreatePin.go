package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/NazarM11/TerestPin/internal/api"
	"github.com/NazarM11/TerestPin/internal/api/middleware"
	"github.com/NazarM11/TerestPin/internal/api/utils"
	"github.com/NazarM11/TerestPin/internal/database"
	"github.com/google/uuid"
)

func CreatePin(cfg *api.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value(middleware.UserIDKey).(uuid.UUID)
		if !ok {
			utils.RespondWithError(w, 401, "Unauthorized", nil)
			return
		}

		err := r.ParseMultipartForm(50 << 20)
		if err != nil {
			utils.RespondWithError(w, 400, "File too large", err)
			return
		}

		file, header, err := r.FormFile("image")
		if err != nil {
			utils.RespondWithError(w, 400, "Invalid image", err)
			return
		}
		defer file.Close()

		fileExt := filepath.Ext(header.Filename)
		allowedExts := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".gif": true}

		if !allowedExts[strings.ToLower(fileExt)] {
			utils.RespondWithError(w, 400, "Invalid file type", nil)
			return
		}

		newFilename := fmt.Sprintf("%s%s", uuid.New().String(), fileExt)
		fullPath := filepath.Join(cfg.Uploads, newFilename)

		title := strings.TrimSpace(r.FormValue("title"))
		if len(title) > 100 {
			utils.RespondWithError(w, 400, "Title too long", nil)
			return
		} else if len(title) < 3 {
			utils.RespondWithError(w, 400, "Title too short", nil)
			return
		}
		buff := make([]byte, 512)
		if _, err := file.Read(buff); err != nil {
			utils.RespondWithError(w, 400, "Bad file", err)
			return
		}

		if _, err := file.Seek(0, io.SeekStart); err != nil {
			utils.RespondWithError(w, 500, "Internal error", err)
			return
		}

		contentType := http.DetectContentType(buff)
		if !strings.HasPrefix(contentType, "image/") {
			utils.RespondWithError(w, 400, "File is not an image", nil)
			return
		}

		dst, err := os.Create(fullPath)
		if err != nil {
			utils.RespondWithError(w, 500, "Couldn't save file", err)
			return
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			utils.RespondWithError(w, 500, "Failed to save file contents", err)
			return
		}

		publicURL := fmt.Sprintf("/uploads/%s", newFilename)
		now := time.Now().UTC()

		pin, err := cfg.DB.CreatePin(r.Context(), database.CreatePinParams{
			ID:        uuid.New(),
			CreatedAt: now,
			UpdatedAt: now,
			Title:     title,
			ImageUrl:  publicURL,
			UserID:    userID,
		})

		if err != nil {
			utils.RespondWithError(w, 500, "Database error", err)
			return
		}

		utils.RespondWithJSON(w, 201, pin)
	}
}
