package auth

import (
	"net/http"
	"strings"
	"errors"
)

var ErrNoAuthHeaderIncluded = errors.New("no auth header included in request")

func GetBearerToken(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")

	if authHeader == "" {
    	return "", ErrNoAuthHeaderIncluded
	}

	tokenString := strings.Split(authHeader, " ")
	if len(tokenString) < 2 || tokenString[0] != "Bearer" {
		return "", ErrNoAuthHeaderIncluded
	}
	return tokenString[1], nil
}