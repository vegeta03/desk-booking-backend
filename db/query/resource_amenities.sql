-- name: CreateResourceAmenity :exec
INSERT INTO resource_amenities (resource_id, amenity_id)
VALUES ($1, $2) 
RETURNING *;

-- name: GetResourceAmenitiesByResource :many
SELECT * 
FROM resource_amenities
WHERE resource_id = $1;

-- name: DeleteResourceAmenity :exec
DELETE FROM resource_amenities
WHERE resource_id = $1 AND amenity_id = $2;
