package main

import (
	"time"

	"github.com/bsach64/blogAggregator/internal/database"
	"github.com/google/uuid"
)

type Feed struct {
	ID            uuid.UUID  `json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	Name          string     `json:"name"`
	Url           string     `json:"url"`
	UserID        uuid.UUID  `json:"user_id"`
	LastFetchedAt *time.Time `json:"last_fetched_at"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

type Post struct {
	ID          uuid.UUID  `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Title       string     `json:"title"`
	Url         string     `json:"url"`
	Description *string    `json:"description"`
	PublishedAt *time.Time `json:"published_at"`
	FeedID      uuid.UUID  `json:"feed_id"`
}

func userFromDatabaseUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.ApiKey,
	}
}

func feedFromDatabaseFeed(feed database.Feed) Feed {
	f := Feed{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		UserID:    feed.UserID,
		Url:       feed.Url,
	}

	var t *time.Time
	if feed.LastFetchedAt.Valid {
		t = &feed.LastFetchedAt.Time
	} else {
		t = nil
	}
	f.LastFetchedAt = t
	return f
}

func feedfollowFromDatabaseFeedFollow(feedfollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        feedfollow.ID,
		CreatedAt: feedfollow.CreatedAt,
		UpdatedAt: feedfollow.UpdatedAt,
		UserID:    feedfollow.UserID,
		FeedID:    feedfollow.FeedID,
	}

}

func postFromDatabasePost(post database.Post) Post {
	p := Post{
		ID:        post.ID,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		Title:     post.Title,
		Url:       post.Url,
		FeedID:    post.FeedID,
	}
	var t *time.Time
	var descrip *string
	if post.Description.Valid {
		descrip = &post.Description.String
	}
	if post.PublishedAt.Valid {
		t = &post.PublishedAt.Time
	}
	p.Description = descrip
	p.PublishedAt = t
	return p
}
