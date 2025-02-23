-- name: CreateBookingStatus :one
INSERT INTO booking_status (status, description)
VALUES ($1, $2)
RETURNING *;

-- name: GetBookingStatus :one
SELECT * 
FROM booking_status
WHERE status = $1 LIMIT 1;

-- name: ListBookingStatuses :many
SELECT * 
FROM booking_status
ORDER BY status;

-- name: UpdateBookingStatus :exec
UPDATE booking_status
SET description = $2
WHERE status = $1;

-- name: DeleteBookingStatus :exec
DELETE FROM booking_status
WHERE status = $1;
