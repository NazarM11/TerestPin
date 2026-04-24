package middleware

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/NazarM11/TerestPin/internal/api"
	"github.com/NazarM11/TerestPin/internal/api/utils"
	"github.com/NazarM11/TerestPin/internal/auth"
)

type contextKey string
const UserIDKey contextKey = "userID"

func MiddlewareAuth(cfg *api.ApiConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString, err := auth.GetBearerToken(r.Header)
			if err != nil {
				utils.RespondWithError(w, 401, "Unauthorized", err)
				return
			}

			userID, err := auth.ValidateJWT(tokenString, cfg.JwtSecret)
			if err != nil {
				utils.RespondWithError(w, 401, "Invalid token", err)
				return
			}

			ctx := context.WithValue(r.Context(), UserIDKey, userID)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func MiddlewareLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf(
			"[%s] %s %s",
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	})
}