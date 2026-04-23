-- name: CreatePin :one
INSERT INTO pins (id, created_at, updated_at, image_url, title, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetPin :one
SELECT * FROM pins WHERE id = $1;

-- name: DeletePin :exec
DELETE FROM pins WHERE id = $1;