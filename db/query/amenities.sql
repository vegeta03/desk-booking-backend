-- name: CreateAmenity :one
INSERT INTO amenities (id, name, description)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetAmenityByID :one
SELECT * 
FROM amenities
WHERE id = $1 LIMIT 1;

-- name: ListAmenities :many
SELECT * 
FROM amenities
ORDER BY name;

-- name: UpdateAmenity :exec
UPDATE amenities
SET name = $2, description = $3
WHERE id = $1;

-- name: DeleteAmenity :exec
DELETE FROM amenities
WHERE id = $1;
