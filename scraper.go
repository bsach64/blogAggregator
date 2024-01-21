package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/bsach64/blogAggregator/internal/database"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func startScraping(
	db *database.Queries,
	concurrency int,
	timeBetweenRequest time.Duration,
) {
	log.Printf("Scraping on %v go runtines every %v seconds\n", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Println("error fetching feeds: ", err)
			continue
		}
		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()
	_, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println(err)
		return
	}

	rss, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println(err)
		return
	}
	for _, entry := range rss.Channel.Item {
		postParams := database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title:     entry.Title,
			Url:       entry.Link,
			FeedID:    feed.ID,
		}
		if entry.Description != "" {
			postParams.Description = sql.NullString{
				String: entry.Description,
				Valid:  true,
			}
		}
		t, err := time.Parse(time.RFC1123Z, entry.PubDate)
		if err != nil {
			log.Println("Could not parse pubDate:", err)
		}
		postParams.PublishedAt = sql.NullTime{Time: t, Valid: true}
		_, err = db.CreatePost(context.Background(), postParams)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Println(err.Error())
		}
	}
	log.Printf("Scraped: %v with %v posts..\n", rss.Channel.Title, len(rss.Channel.Item))
}
