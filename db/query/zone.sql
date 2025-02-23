-- name: CreateZone :one
INSERT INTO zones (id, floor_id, department_id, name, description)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetZoneByID :one
SELECT * 
FROM zones
WHERE id = $1 LIMIT 1;

-- name: ListZonesByFloor :many
SELECT * 
FROM zones
WHERE floor_id = $1
ORDER BY name;

-- name: UpdateZone :exec
UPDATE zones
SET floor_id = $2, department_id = $3, name = $4, description = $5
WHERE id = $1;

-- name: DeleteZone :exec
DELETE FROM zones
WHERE id = $1;
