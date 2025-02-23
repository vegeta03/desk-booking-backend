-- name: CreateUserRole :exec
INSERT INTO user_roles (user_email, role)
VALUES ($1, $2) 
RETURNING *;

-- name: GetUserRoles :many
SELECT * 
FROM user_roles
WHERE user_email = $1;

-- name: DeleteUserRole :exec
DELETE FROM user_roles
WHERE user_email = $1 AND role = $2;
