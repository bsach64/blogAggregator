package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/bsach64/blogAggregator/internal/database"
	"github.com/go-chi/chi/v5"
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
		feedfollowFromDatabaseFeedFollow(feedFollow),
	)
}

func (cfg *apiConfig) handleDeleteFeedFollow(w http.ResponseWriter, req *http.Request, user database.User) {
	feedIDStr := chi.URLParam(req, "feedfollowID")
	
	feedID, err := uuid.Parse(feedIDStr)
	if err != nil {
		respondWithError(
			w,
			http.StatusBadRequest,
			err.Error(),
		)
	}

	deleteParms := database.DeleteFeedFollowParams{
		UserID: user.ID,
		ID: feedID,
	}

	err = cfg.DB.DeleteFeedFollow(req.Context(), deleteParms)
	if err != nil {
		respondWithError(
			w,
			http.StatusBadRequest,
			err.Error(),
		)
	}
	respondWithJSON(w, http.StatusOK, struct{}{})
}


func (cfg *apiConfig) handleGetFeedFollows(w http.ResponseWriter, req *http.Request, user database.User) {
	feeds, err := cfg.DB.AllFeedFollows(req.Context(), user.ID)
	if err != nil {
		respondWithError(
			w,
			http.StatusBadRequest,
			err.Error(),
		)
	}

	respondWithJSON(
		w,
		http.StatusOK,
		feeds,
	)
}
