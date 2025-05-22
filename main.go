package main

import (
	"context"
	"log"
	"simple_bank/api"
	db "simple_bank/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbSource      = "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("can not start:", err)
	}
}
