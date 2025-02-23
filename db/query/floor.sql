-- name: CreateFloor :one
INSERT INTO floors (floor_number, floor_plan, building_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetFloorByID :one
SELECT * 
FROM floors
WHERE id = $1 LIMIT 1;

-- name: ListFloorsByBuilding :many
SELECT * 
FROM floors
WHERE building_id = $1
ORDER BY floor_number;

-- name: UpdateFloor :exec
UPDATE floors
SET floor_number = $2, floor_plan = $3, building_id = $4
WHERE id = $1;

-- name: DeleteFloor :exec
DELETE FROM floors
WHERE id = $1;
