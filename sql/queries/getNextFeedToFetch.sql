-- name: GetNextFeedToFetch :one

SELECT id, name, url, last_fetched_at, created_at, updated_at
FROM feeds
ORDER BY last_fetched_at NULLS FIRST, updated_at
LIMIT 1;
