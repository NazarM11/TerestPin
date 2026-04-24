-- name: CreatePin :one
INSERT INTO pins (id, created_at, updated_at, image_url, title, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetPin :one
SELECT * FROM pins WHERE id = $1;

-- name: GetPins :many
SELECT * FROM pins
WHERE title ILIKE $3
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: DeletePin :exec
DELETE FROM pins WHERE id = $1 AND user_id = $2;

-- name: GetPinsByUserID :many
SELECT * FROM pins WHERE user_id = $1 ORDER BY created_at DESC;