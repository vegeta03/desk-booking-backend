-- name: CreateResourceStatus :one
INSERT INTO resource_status (status, description)
VALUES ($1, $2)
RETURNING *;

-- name: GetResourceStatus :one
SELECT * 
FROM resource_status
WHERE status = $1 LIMIT 1;

-- name: ListResourceStatuses :many
SELECT * 
FROM resource_status
ORDER BY status;

-- name: UpdateResourceStatus :exec
UPDATE resource_status
SET description = $2
WHERE status = $1;

-- name: DeleteResourceStatus :exec
DELETE FROM resource_status
WHERE status = $1;
