// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: role_permissions.sql

package db

import (
	"context"
)

const createRolePermission = `-- name: CreateRolePermission :exec
INSERT INTO role_permissions (role_id, permission_id)
VALUES ($1, $2) 
RETURNING role_id, permission_id
`

type CreateRolePermissionParams struct {
	RoleID       string `json:"role_id"`
	PermissionID string `json:"permission_id"`
}

func (q *Queries) CreateRolePermission(ctx context.Context, arg CreateRolePermissionParams) error {
	_, err := q.db.Exec(ctx, createRolePermission, arg.RoleID, arg.PermissionID)
	return err
}

const deleteRolePermission = `-- name: DeleteRolePermission :exec
DELETE FROM role_permissions
WHERE role_id = $1 AND permission_id = $2
`

type DeleteRolePermissionParams struct {
	RoleID       string `json:"role_id"`
	PermissionID string `json:"permission_id"`
}

func (q *Queries) DeleteRolePermission(ctx context.Context, arg DeleteRolePermissionParams) error {
	_, err := q.db.Exec(ctx, deleteRolePermission, arg.RoleID, arg.PermissionID)
	return err
}

const getPermissionsByRole = `-- name: GetPermissionsByRole :many
SELECT role_id, permission_id 
FROM role_permissions
WHERE role_id = $1
`

func (q *Queries) GetPermissionsByRole(ctx context.Context, roleID string) ([]RolePermission, error) {
	rows, err := q.db.Query(ctx, getPermissionsByRole, roleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RolePermission
	for rows.Next() {
		var i RolePermission
		if err := rows.Scan(&i.RoleID, &i.PermissionID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
