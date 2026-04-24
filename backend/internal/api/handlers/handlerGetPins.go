package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/NazarM11/TerestPin/internal/api"
	"github.com/NazarM11/TerestPin/internal/api/utils"
	"github.com/NazarM11/TerestPin/internal/database"
	"github.com/google/uuid"
)

type PinResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `json:"title"`
	ImageUrl  string    `json:"image_url"`
	UserID    uuid.UUID `json:"user_id"`
}

func GetPins(cfg *api.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit := 20
		pageInt := 1
		page := r.URL.Query().Get("page")
		if page != "" {
			var err error
			pageInt, err = strconv.Atoi(page)
			if err != nil {
				utils.RespondWithError(w, 500, "Couldnt get page number", err)
				return
			}
			if pageInt <= 1 {
				pageInt = 1
			}
		}
		searchQuery := r.URL.Query().Get("search")
		search := fmt.Sprintf("%%%s%%", searchQuery)

		offset := (pageInt - 1) * limit

		dbPins, err := cfg.DB.GetPins(r.Context(), database.GetPinsParams{
			Limit:  int32(limit),
			Offset: int32(offset),
			Title:  search,
		})
		if err != nil {
			utils.RespondWithError(w, 500, "Couldn't retrieve pins", err)
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
