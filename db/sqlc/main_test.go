package db

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib" // pgx driver
)

const (
	dbDriver = "pgx"
	dbSource = "postgresql://root:Batman@localhost:5432/desk_booking?sslmode=disable"
)

var testQueries *Queries

func TestMain(mm *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}

	testQueries = New(conn)
}
