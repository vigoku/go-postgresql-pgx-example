package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/georgysavva/scany/pgxscan"
)

type User struct {
	ID    string
	Name  string
	Email string
	Age   int
}

func main() {
	ctx := context.Background()
	db, _ := pgxpool.Connect(ctx, "postgres://postgres:password@localhost:5432/postgres")

	var users []*User
	pgxscan.Select(ctx, db, &users, `SELECT id, name, email, age FROM users`)
	// users variable now contains data from all rows.

	for i := 0; i < len(users); i++ {
		log.Print(i)
		log.Println(users[i].ID, users[i].Name, users[i].Email, users[i].Age)
	}
}
