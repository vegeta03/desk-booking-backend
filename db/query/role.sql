-- name: CreateRole :one
INSERT INTO roles (role, description)
VALUES ($1, $2)
RETURNING *;

-- name: GetRole :one
SELECT * 
FROM roles
WHERE role = $1 LIMIT 1;

-- name: ListRoles :many
SELECT * 
FROM roles
ORDER BY role;

-- name: UpdateRole :exec
UPDATE roles
SET description = $2
WHERE role = $1;

-- name: DeleteRole :exec
DELETE FROM roles
WHERE role = $1;
