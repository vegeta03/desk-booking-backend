-- name: CreatePermission :one
INSERT INTO permissions (permission, description)
VALUES ($1, $2)
RETURNING *;

-- name: GetPermission :one
SELECT * 
FROM permissions
WHERE permission = $1 LIMIT 1;

-- name: ListPermissions :many
SELECT * 
FROM permissions
ORDER BY permission;

-- name: UpdatePermission :exec
UPDATE permissions
SET description = $2
WHERE permission = $1;

-- name: DeletePermission :exec
DELETE FROM permissions
WHERE permission = $1;
