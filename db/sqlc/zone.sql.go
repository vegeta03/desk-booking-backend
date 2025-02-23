// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: zone.sql

package db

import (
	"context"
)

const createZone = `-- name: CreateZone :one
INSERT INTO zones (id, floor_id, department_id, name, description)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, floor_id, department_id, name, description
`

type CreateZoneParams struct {
	ID           string `json:"id"`
	FloorID      int32  `json:"floor_id"`
	DepartmentID int32  `json:"department_id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
}

func (q *Queries) CreateZone(ctx context.Context, arg CreateZoneParams) (Zone, error) {
	row := q.db.QueryRow(ctx, createZone,
		arg.ID,
		arg.FloorID,
		arg.DepartmentID,
		arg.Name,
		arg.Description,
	)
	var i Zone
	err := row.Scan(
		&i.ID,
		&i.FloorID,
		&i.DepartmentID,
		&i.Name,
		&i.Description,
	)
	return i, err
}

const deleteZone = `-- name: DeleteZone :exec
DELETE FROM zones
WHERE id = $1
`

func (q *Queries) DeleteZone(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteZone, id)
	return err
}

const getZoneByID = `-- name: GetZoneByID :one
SELECT id, floor_id, department_id, name, description 
FROM zones
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetZoneByID(ctx context.Context, id string) (Zone, error) {
	row := q.db.QueryRow(ctx, getZoneByID, id)
	var i Zone
	err := row.Scan(
		&i.ID,
		&i.FloorID,
		&i.DepartmentID,
		&i.Name,
		&i.Description,
	)
	return i, err
}

const listZonesByFloor = `-- name: ListZonesByFloor :many
SELECT id, floor_id, department_id, name, description 
FROM zones
WHERE floor_id = $1
ORDER BY name
`

func (q *Queries) ListZonesByFloor(ctx context.Context, floorID int32) ([]Zone, error) {
	rows, err := q.db.Query(ctx, listZonesByFloor, floorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Zone
	for rows.Next() {
		var i Zone
		if err := rows.Scan(
			&i.ID,
			&i.FloorID,
			&i.DepartmentID,
			&i.Name,
			&i.Description,
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

const updateZone = `-- name: UpdateZone :exec
UPDATE zones
SET floor_id = $2, department_id = $3, name = $4, description = $5
WHERE id = $1
`

type UpdateZoneParams struct {
	ID           string `json:"id"`
	FloorID      int32  `json:"floor_id"`
	DepartmentID int32  `json:"department_id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
}

func (q *Queries) UpdateZone(ctx context.Context, arg UpdateZoneParams) error {
	_, err := q.db.Exec(ctx, updateZone,
		arg.ID,
		arg.FloorID,
		arg.DepartmentID,
		arg.Name,
		arg.Description,
	)
	return err
}
