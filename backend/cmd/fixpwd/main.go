package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	ctx := context.Background()

	// Connect directly to the database
	connString := "postgres://postgres:sulcaprado@localhost:5433/unsch_horarios?sslmode=disable"
	db, err := pgxpool.New(ctx, connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer db.Close()

	// Generate hash for admin123
	hash, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	fmt.Printf("Generated hash: %s\n", string(hash))

	// Update the password
	result, err := db.Exec(ctx,
		"UPDATE usuario SET password_hash = $1 WHERE email = $2",
		string(hash), "admin@unsch.edu.pe",
	)
	if err != nil {
		log.Fatalf("Failed to update password: %v", err)
	}

	fmt.Printf("Rows affected: %d\n", result.RowsAffected())
	fmt.Println("Password updated successfully!")
}
