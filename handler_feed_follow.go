package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/quangrau/rssagg/internal/database"
	"net/http"
	"time"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(rw http.ResponseWriter, r *http.Request, user database.User) {
	type reqParams struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := reqParams{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(rw, http.StatusBadRequest, fmt.Sprintf("Decode Request Body Error: %v", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		respondWithError(rw, http.StatusInternalServerError, fmt.Sprintf("Create Feed Follow Error: %v", err))
	}

	respondWithJSON(rw, http.StatusOK, databaseFeedFollowToFeedFollow(feedFollow))
}
