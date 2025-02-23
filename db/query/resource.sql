-- name: CreateResource :one
INSERT INTO resources (floor_id, zone_id, resource_type, identifier, capacity, status)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetResourceByID :one
SELECT * 
FROM resources
WHERE id = $1 LIMIT 1;

-- name: ListResourcesByZone :many
SELECT * 
FROM resources
WHERE zone_id = $1
ORDER BY identifier;

-- name: UpdateResource :exec
UPDATE resources
SET floor_id = $2, zone_id = $3, resource_type = $4, identifier = $5, capacity = $6, status = $7
WHERE id = $1;

-- name: DeleteResource :exec
DELETE FROM resources
WHERE id = $1;
