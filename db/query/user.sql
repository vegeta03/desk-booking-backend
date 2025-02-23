-- name: CreateUser :one
INSERT INTO users (email, name, employee_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * 
FROM users
WHERE email = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * 
FROM users
ORDER BY email;

-- name: UpdateUser :exec
UPDATE users
SET name = $2, employee_id = $3
WHERE email = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE email = $1;
