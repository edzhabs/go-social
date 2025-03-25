package main

import (
	"log"

	"github.com/edzhabs/social/internal/db"
	"github.com/edzhabs/social/internal/env"
	"github.com/edzhabs/social/internal/store"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	addr := env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/socialnetwork?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)
	db.Seed(store)
}
