-- name: CreateEmployee :one
INSERT INTO employees (employee_id, department_id, team_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetEmployeeByID :one
SELECT * 
FROM employees
WHERE employee_id = $1 LIMIT 1;

-- name: ListEmployees :many
SELECT * 
FROM employees
ORDER BY employee_id;

-- name: UpdateEmployee :exec
UPDATE employees
SET department_id = $2, team_id = $3
WHERE employee_id = $1;

-- name: DeleteEmployee :exec
DELETE FROM employees
WHERE employee_id = $1;
