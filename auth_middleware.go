package main

import (
	"net/http"
	"strings"

	"github.com/bsach64/blogAggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := strings.TrimPrefix(r.Header.Get("Authorization"), "ApiKey ")
		data, err := cfg.DB.GetUserApi(r.Context(), auth)
		if err != nil {
			respondWithError(
				w,
				http.StatusBadRequest,
				"User Could not be found",
			)
			return
		}
		handler(w, r, data)
		return 
	})	
}
