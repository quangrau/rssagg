package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts an API Key from the headers of an HTTP request.
// Example:
// Authorization: ApiKey {YOUR-API-KEY-HERE}
func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if len(authHeader) == 0 {
		return "", errors.New("Authorization header not found")
	}

	authHeaderVals := strings.Split(authHeader, " ")
	if len(authHeaderVals) != 2 {
		return "", errors.New("Authorization header format is invalid")
	}

	if authHeaderVals[0] != "ApiKey" {
		return "", errors.New("Authorization header format is invalid")
	}

	return authHeaderVals[1], nil
}
