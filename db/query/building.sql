-- name: CreateBuilding :one
INSERT INTO buildings (name, description)
VALUES ($1, $2)
RETURNING *;

-- name: GetBuildingByID :one
SELECT * 
FROM buildings
WHERE id = $1 LIMIT 1;

-- name: ListBuildings :many
SELECT * 
FROM buildings
ORDER BY name;

-- name: UpdateBuilding :exec
UPDATE buildings
SET name = $2, description = $3
WHERE id = $1;

-- name: DeleteBuilding :exec
DELETE FROM buildings
WHERE id = $1;
