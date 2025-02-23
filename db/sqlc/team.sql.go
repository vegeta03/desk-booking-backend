// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: team.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createTeam = `-- name: CreateTeam :one
INSERT INTO teams (name, description, department_id)
VALUES ($1, $2, $3)
RETURNING id, name, description, department_id
`

type CreateTeamParams struct {
	Name         string      `json:"name"`
	Description  string      `json:"description"`
	DepartmentID pgtype.Int4 `json:"department_id"`
}

func (q *Queries) CreateTeam(ctx context.Context, arg CreateTeamParams) (Team, error) {
	row := q.db.QueryRow(ctx, createTeam, arg.Name, arg.Description, arg.DepartmentID)
	var i Team
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.DepartmentID,
	)
	return i, err
}

const deleteTeam = `-- name: DeleteTeam :exec
DELETE FROM teams
WHERE id = $1
`

func (q *Queries) DeleteTeam(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteTeam, id)
	return err
}

const getTeamByID = `-- name: GetTeamByID :one
SELECT id, name, description, department_id 
FROM teams
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTeamByID(ctx context.Context, id int32) (Team, error) {
	row := q.db.QueryRow(ctx, getTeamByID, id)
	var i Team
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.DepartmentID,
	)
	return i, err
}

const listTeamsByDepartment = `-- name: ListTeamsByDepartment :many
SELECT id, name, description, department_id 
FROM teams
WHERE department_id = $1
ORDER BY name
`

func (q *Queries) ListTeamsByDepartment(ctx context.Context, departmentID pgtype.Int4) ([]Team, error) {
	rows, err := q.db.Query(ctx, listTeamsByDepartment, departmentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Team
	for rows.Next() {
		var i Team
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.DepartmentID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTeam = `-- name: UpdateTeam :exec
UPDATE teams
SET name = $2, description = $3, department_id = $4
WHERE id = $1
`

type UpdateTeamParams struct {
	ID           int32       `json:"id"`
	Name         string      `json:"name"`
	Description  string      `json:"description"`
	DepartmentID pgtype.Int4 `json:"department_id"`
}

func (q *Queries) UpdateTeam(ctx context.Context, arg UpdateTeamParams) error {
	_, err := q.db.Exec(ctx, updateTeam,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.DepartmentID,
	)
	return err
}
