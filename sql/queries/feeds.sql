-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetAllFeeds :many
SELECT * FROM feeds;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds ORDER BY last_fetched_at DESC;

-- name: MarkFeedFetched :exec
UPDATE feeds 
SET updated_at = CURRENT_TIME, last_fetched_at = CURRENT_TIME
WHERE id = $1;
