// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: floor.sql

package db

import (
	"context"
)

const createFloor = `-- name: CreateFloor :one
INSERT INTO floors (floor_number, floor_plan, building_id)
VALUES ($1, $2, $3)
RETURNING id, floor_number, floor_plan, building_id
`

type CreateFloorParams struct {
	FloorNumber int32  `json:"floor_number"`
	FloorPlan   string `json:"floor_plan"`
	BuildingID  int32  `json:"building_id"`
}

func (q *Queries) CreateFloor(ctx context.Context, arg CreateFloorParams) (Floor, error) {
	row := q.db.QueryRow(ctx, createFloor, arg.FloorNumber, arg.FloorPlan, arg.BuildingID)
	var i Floor
	err := row.Scan(
		&i.ID,
		&i.FloorNumber,
		&i.FloorPlan,
		&i.BuildingID,
	)
	return i, err
}

const deleteFloor = `-- name: DeleteFloor :exec
DELETE FROM floors
WHERE id = $1
`

func (q *Queries) DeleteFloor(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteFloor, id)
	return err
}

const getFloorByID = `-- name: GetFloorByID :one
SELECT id, floor_number, floor_plan, building_id 
FROM floors
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetFloorByID(ctx context.Context, id int32) (Floor, error) {
	row := q.db.QueryRow(ctx, getFloorByID, id)
	var i Floor
	err := row.Scan(
		&i.ID,
		&i.FloorNumber,
		&i.FloorPlan,
		&i.BuildingID,
	)
	return i, err
}

const listFloorsByBuilding = `-- name: ListFloorsByBuilding :many
SELECT id, floor_number, floor_plan, building_id 
FROM floors
WHERE building_id = $1
ORDER BY floor_number
`

func (q *Queries) ListFloorsByBuilding(ctx context.Context, buildingID int32) ([]Floor, error) {
	rows, err := q.db.Query(ctx, listFloorsByBuilding, buildingID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Floor
	for rows.Next() {
		var i Floor
		if err := rows.Scan(
			&i.ID,
			&i.FloorNumber,
			&i.FloorPlan,
			&i.BuildingID,
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

const updateFloor = `-- name: UpdateFloor :exec
UPDATE floors
SET floor_number = $2, floor_plan = $3, building_id = $4
WHERE id = $1
`

type UpdateFloorParams struct {
	ID          int32  `json:"id"`
	FloorNumber int32  `json:"floor_number"`
	FloorPlan   string `json:"floor_plan"`
	BuildingID  int32  `json:"building_id"`
}

func (q *Queries) UpdateFloor(ctx context.Context, arg UpdateFloorParams) error {
	_, err := q.db.Exec(ctx, updateFloor,
		arg.ID,
		arg.FloorNumber,
		arg.FloorPlan,
		arg.BuildingID,
	)
	return err
}
