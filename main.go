package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"

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
	databaseUrl := "postgres://postgres:password@localhost:5432/postgres"
	db, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close(context.Background())
	//GET ALL ROWS
	// var users []*User
	// pgxscan.Select(ctx, db, &users, `SELECT id, name, email, age FROM users`)
	// // users variable now contains data from all rows.

	// for i := 0; i < len(users); i++ {
	// 	log.Print(i)
	// 	log.Println(users[i].ID, users[i].Name, users[i].Email, users[i].Age)
	rows, _ := db.Query(ctx, `SELECT id, name, email, age FROM users`)
	defer rows.Close()
	for rows.Next() {
		var user User
		if err := pgxscan.ScanRow(&user, rows); err != nil {
			// Handle row scanning error.
			fmt.Fprintf(os.Stderr, "End of rows: %v\n", err)
			return
		}
		// user variable now contains data from the current row.
		log.Println(user.ID, user.Name, user.Email, user.Age)
	}
	if err := rows.Err(); err != nil {
		// Handle rows final error.
		fmt.Fprintf(os.Stderr, "End of rows: %v\n", err)
	}
}
