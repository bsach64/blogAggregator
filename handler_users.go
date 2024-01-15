package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/bsach64/blogAggregator/internal/database"
	"github.com/google/uuid"
)

func (api *apiConfig) handleCreateUsers(w http.ResponseWriter, req *http.Request) {
	type paramaters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(req.Body)
	params := paramaters{}
	err := decoder.Decode(&params)

	if err != nil {
		log.Println(err)
		handleErr(w, req)
		return
	}

	newUser := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
	}

	user, err := api.DB.CreateUser(req.Context(), newUser)
	if err != nil {
		log.Println(err)
		handleErr(w, req)
		return
	}
	respondWithJSON(
		w,
		http.StatusCreated,
		user,
	)

}

func (api *apiConfig) handleGetUserFromApi(w http.ResponseWriter, req *http.Request, user database.User) {
	respondWithJSON(
		w,
		http.StatusOK,
		user,
	)
}
