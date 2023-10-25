package main

import (
	"log"

	"github.com/quocbang/versioned-migration/connection"
	"github.com/quocbang/versioned-migration/migrations"
)

func main() {
	db, err := connection.NewDatabaseConn(connection.MySQLConfig{
		Address:  "localhost",
		Port:     3306,
		Name:     "example",
		Username: "root",
		Password: "123",
	})
	if err != nil {
		log.Fatalf("failed to connect database, error: %v", err)
	}

	if err := migrations.Up(db); err != nil {
		log.Fatalf("failed to migrate, error: %v", err)
	}
}
