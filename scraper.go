package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/bsach64/blogAggregator/internal/database"
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
		log.Println(rss.Channel.Title ,entry.Title)
	}
	log.Println("Done: ", rss.Channel.Title)
}
