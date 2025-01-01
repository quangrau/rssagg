package main

import (
	"context"
	"github.com/quangrau/rssagg/internal/database"
	"log"
	"sync"
	"time"
)

func runScraperWorker(db *database.Queries, concurrency int, duration time.Duration) {
	log.Printf("Scraper worker started on %v goroutines every %s", concurrency, duration)

	ticker := time.NewTicker(duration)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Printf("Error fetching feeds: %v", err)
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

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Error updating feed %v: %v", feed.ID, err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Printf("Error fetching feed %v: %v", feed.ID, err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		log.Printf("Found post: %v", item.Title)
	}
	log.Printf("-- Feed %s collected, found %v post --", feed.Name, len(rssFeed.Channel.Item))
}
