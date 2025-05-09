-- name: LookUpFeedByUrl :one

SELECT * FROM feeds
WHERE url = $1;
