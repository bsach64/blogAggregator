package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/bsach64/blogAggregator/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Could not get environment variables..")
	}
	port := os.Getenv("PORT")
	dbURL := os.Getenv("DB")

	db, err := sql.Open("postgres", dbURL)

	api := apiConfig{
		DB: database.New(db),
	}
	
	router := chi.NewRouter()
	handler := cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	})
	router.Use(handler)
	routerV1 := chi.NewRouter()

	routerV1.Get(
		"/readiness",
		handleReadiness,
	)

	routerV1.Get(
		"/err",
		handleErr,
	)

	routerV1.Post(
		"/users",
		api.handleCreateUsers,
	)

	routerV1.Get(
		"/users",
		api.middlewareAuth(api.handleGetUserFromApi),
	)

	routerV1.Post(
		"/feeds",
		api.middlewareAuth(api.handleCreateFeed),
	)
	
	routerV1.Get(
		"/all_feeds",
		api.handleGetAllFeeds,
	)
	
	routerV1.Post(
		"/feed_follows",
		api.middlewareAuth(api.handleFeedFollow),
	)

	router.Mount("/v1", routerV1)
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	err = server.ListenAndServe()
	log.Print(err)
}
