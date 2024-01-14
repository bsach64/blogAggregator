package main

import "net/http"

func handleReadiness(w http.ResponseWriter, req *http.Request) {
	respondWithJSON(
		w,
		http.StatusOK,
		struct {
			Status string `json:"status"`
		}{
			Status: "ok",
		},
	)
}
