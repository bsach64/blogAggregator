package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)


func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Could not get environment variables..")	
	}
	port := os.Getenv("PORT")
	r := chi.NewRouter()
	handler := cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	})
	r.Use(handler)
	rV1 := chi.NewRouter()
	r.Mount("/v1", rV1)
	server := &http.Server{
		Addr: ":" + port,
		Handler: r,
	}
	err = server.ListenAndServe()
	log.Print(err)
}
