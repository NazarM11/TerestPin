package auth

import (
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")

	if authHeader == "" {
    	return "", ErrNoAuthHeaderIncluded
	}

	api_key := strings.Split(authHeader, " ")
	if len(api_key) < 2 || api_key[0] != "ApiKey" {
		return "", ErrNoAuthHeaderIncluded
	}
	return api_key[1], nil
}