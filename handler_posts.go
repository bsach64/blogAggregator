package main

import (
	"net/http"
	"strconv"

	"github.com/bsach64/blogAggregator/internal/database"
)

func (cfg *apiConfig) handleGetPosts(w http.ResponseWriter, req *http.Request, user database.User) {
	queries := req.URL.Query()
	
	dbPosts, err := cfg.DB.GetPostsbyUser(req.Context(), user.ID)
	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			err.Error(),
		)
	}
	posts := make([]Post, 0)
	for _, p := range dbPosts {
		posts = append(posts, postFromDatabasePost(p))
	}
	
	limit, err := strconv.Atoi(queries.Get("limit"))
	if err != nil {
		limit = len(posts)	
	}
	respondWithJSON(
		w,
		http.StatusOK,
		posts[:limit],
	)
}
