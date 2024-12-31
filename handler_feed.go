package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/quangrau/rssagg/internal/database"
	"net/http"
	"time"
)

func (apiCfg *apiConfig) handlerCreateFeed(rw http.ResponseWriter, r *http.Request, user database.User) {
	type reqParams struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := reqParams{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(rw, http.StatusBadRequest, fmt.Sprintf("Decode Request Body Error: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(rw, http.StatusInternalServerError, fmt.Sprintf("Create Feed Error: %v", err))
		return
	}

	respondWithJSON(rw, http.StatusOK, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeeds(rw http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(rw, http.StatusInternalServerError, fmt.Sprintf("Get Feed Error: %v", err))
		return
	}

	respondWithJSON(rw, http.StatusOK, databaseFeedsToFeeds(feeds))
}
