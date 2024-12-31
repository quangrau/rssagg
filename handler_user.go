package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/quangrau/rssagg/internal/database"
	"net/http"
	"time"
)

func (apiCfg *apiConfig) handlerCreateUser(rw http.ResponseWriter, r *http.Request) {
	type reqParams struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := reqParams{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(rw, http.StatusBadRequest, fmt.Sprintf("Decode Request Body Error: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(rw, http.StatusInternalServerError, fmt.Sprintf("Create User Error: %v", err))
	}

	respondWithJSON(rw, http.StatusOK, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(rw http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(rw, http.StatusOK, databaseUserToUser(user))
}
