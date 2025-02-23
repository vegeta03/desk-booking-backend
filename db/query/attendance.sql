-- name: CreateAttendance :one
INSERT INTO attendances (employee_id, booking_id, check_in_time, check_out_time, duration)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetAttendanceByID :one
SELECT * 
FROM attendances
WHERE id = $1 LIMIT 1;

-- name: ListAttendancesByEmployee :many
SELECT * 
FROM attendances
WHERE employee_id = $1
ORDER BY check_in_time;

-- name: UpdateAttendance :exec
UPDATE attendances
SET employee_id = $2, booking_id = $3, check_in_time = $4, check_out_time = $5, duration = $6
WHERE id = $1;

-- name: DeleteAttendance :exec
DELETE FROM attendances
WHERE id = $1;
