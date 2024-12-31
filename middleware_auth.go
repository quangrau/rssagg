package main

import (
	"fmt"
	"github.com/quangrau/rssagg/internal/auth"
	"github.com/quangrau/rssagg/internal/database"
	"net/http"
)

type authHandler func(w http.ResponseWriter, r *http.Request, user database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(rw, http.StatusUnauthorized, fmt.Sprintf("Auth Error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(rw, http.StatusInternalServerError, fmt.Sprintf("GetUser Error: %v", err))
			return
		}

		handler(rw, r, user)
	}
}
