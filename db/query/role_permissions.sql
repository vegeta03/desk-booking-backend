-- name: CreateRolePermission :exec
INSERT INTO role_permissions (role_id, permission_id)
VALUES ($1, $2) 
RETURNING *;

-- name: GetPermissionsByRole :many
SELECT * 
FROM role_permissions
WHERE role_id = $1;

-- name: DeleteRolePermission :exec
DELETE FROM role_permissions
WHERE role_id = $1 AND permission_id = $2;
