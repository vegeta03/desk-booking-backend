-- name: CreateDepartment :one
INSERT INTO departments (name, description)
VALUES ($1, $2)
RETURNING *;

-- name: GetDepartmentByID :one
SELECT * 
FROM departments
WHERE id = $1 LIMIT 1;

-- name: ListDepartments :many
SELECT * 
FROM departments
ORDER BY name;

-- name: UpdateDepartment :exec
UPDATE departments
SET name = $2, description = $3
WHERE id = $1;

-- name: DeleteDepartment :exec
DELETE FROM departments
WHERE id = $1;
