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
		Url string `json:"url"`
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
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		Name: params.Name,
		Url: params.Url,
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
	
	respondWithJSON(
		w,
		http.StatusCreated,
		feedFromDatabaseFeed(feedEntry),
	)	

}
