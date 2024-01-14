package main

import "net/http"

func handleErr(w http.ResponseWriter, req *http.Request) {
	respondWithError(
		w,
		http.StatusInternalServerError,
		"Internal Server Error",
	)
}
