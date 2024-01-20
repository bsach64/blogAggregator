package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bsach64/blogAggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handleCreateFeed(w http.ResponseWriter, req *http.Request, user database.User) {
	type paramaters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decoder := json.NewDecoder(req.Body)
	params := paramaters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(
			w,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	feed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		Name:      params.Name,
		Url:       params.Url,
	}

	feedEntry, err := cfg.DB.CreateFeed(req.Context(), feed)

	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	type responseStruct struct {
		NewFeed       Feed       `json:"feed"`
		NewFeedFollow FeedFollow `json:"feed_follow"`
	}

	FeedFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feedEntry.ID,
	}

	NewFeedFollow, err := cfg.DB.CreateFeedFollow(req.Context(), FeedFollow)

	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	respondWithJSON(
		w,
		http.StatusCreated,
		responseStruct{
			feedFromDatabaseFeed(feedEntry),
			feedfollowFromDatabaseFeedFollow(NewFeedFollow),
		},
	)

}

func (cfg *apiConfig) handleGetAllFeeds(w http.ResponseWriter, req *http.Request) {
	feeds, err := cfg.DB.GetAllFeeds(req.Context())
	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	fs := make([]Feed, 0)
	for _, f := range feeds {
		fs = append(fs, feedFromDatabaseFeed(f))
	}

	respondWithJSON(
		w,
		http.StatusOK,
		fs,
	)
}
