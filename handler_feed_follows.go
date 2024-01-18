package main

import (
	"net/http"
	"time"
	"encoding/json"
	"log"

	"github.com/bsach64/blogAggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handleFeedFollow(w http.ResponseWriter, req *http.Request, user database.User) {
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}
	
	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		log.Println(err)
		handleErr(w, req)
		return
	}

	createFeedFollow := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: params.FeedId,
	}

	feedFollow, err := cfg.DB.CreateFeedFollow(req.Context(), createFeedFollow)
	
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
		FeedFollow(feedFollow),
	)
}
