package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbSource = "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *pgxpool.Pool

// TestMain initializes the database connection and sets up the queries object for testing
func TestMain(m *testing.M) {
	// creating connection database pool
	var err error
	testDB, err = pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}
	defer testDB.Close()

	// creating query object
	testQueries = New(testDB)

	// run test
	os.Exit(m.Run())
}
