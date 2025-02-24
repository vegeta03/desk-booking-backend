package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
)

const (
	dbDriver = "pgx"
	dbSource = "postgresql://root:Batman@localhost:5432/desk_booking?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := pgx.Connect(context.Background(), dbSource)
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}
	defer conn.Close(context.Background())

	testQueries = New(conn)
	os.Exit(m.Run())
}
