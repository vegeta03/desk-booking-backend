-- name: CreateAuditLogEntry :one
INSERT INTO audit_log (table_name, record_id, action, user_email)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetAuditLogEntryByID :one
SELECT * 
FROM audit_log
WHERE id = $1 LIMIT 1;

-- name: ListAuditLogsByUser :many
SELECT * 
FROM audit_log
WHERE user_email = $1
ORDER BY timestamp DESC;

-- name: DeleteAuditLogEntry :exec
DELETE FROM audit_log
WHERE id = $1;
