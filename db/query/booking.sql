-- name: CreateBooking :one
INSERT INTO bookings (resource_id, employee_id, date, start_time, end_time, status)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetBookingByID :one
SELECT * 
FROM bookings
WHERE id = $1 LIMIT 1;

-- name: ListBookingsByDate :many
SELECT * 
FROM bookings
WHERE date = $1
ORDER BY start_time;

-- name: UpdateBooking :exec
UPDATE bookings
SET resource_id = $2, employee_id = $3, date = $4, start_time = $5, end_time = $6, status = $7
WHERE id = $1;

-- name: DeleteBooking :exec
DELETE FROM bookings
WHERE id = $1;
