-- name: CreateTeam :one
INSERT INTO teams (name, description, department_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetTeamByID :one
SELECT * 
FROM teams
WHERE id = $1 LIMIT 1;

-- name: ListTeamsByDepartment :many
SELECT * 
FROM teams
WHERE department_id = $1
ORDER BY name;

-- name: UpdateTeam :exec
UPDATE teams
SET name = $2, description = $3, department_id = $4
WHERE id = $1;

-- name: DeleteTeam :exec
DELETE FROM teams
WHERE id = $1;
