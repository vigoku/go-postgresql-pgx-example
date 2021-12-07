package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
)

func main() {
	log.Println("starting program")
	// get the database connection URL.
	// usually, this is taken as an environment variable as in below commented out code
	// databaseUrl = os.Getenv("DATABASE_URL")
	// for the time being, let's hard code it as follows. change the values as needed.
	databaseUrl := "postgres://postgres:password@localhost:5432/postgres"
	db, err := pgx.Connect(context.Background(), databaseUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	//to close DB pool
	defer db.Close(context.Background())

	//ExecuteSelectQuery(db)
	//ExecuteFunction(db)
	ExecuteMultiRecord(db)
	//ExcecuteMultiInsert(db);
	log.Println("stopping program")
}

type User struct {
	id            string
	first_name    string
	last_name     string
	date_of_birth string
}

func ExecuteMultiRecord(db *pgx.Conn) {
	ctx := context.Background()
	var users []*User
	err := pgxscan.Select(ctx, db, &users, `SELECT id, first_name, last_name, date_of_birth FROM users`)
	if err != nil {
		return
	}

	for (int i = 0; i < users.Len(); i++) {
		log.Print(i);
		log.Print(users[i].id, users[i].first_name, )
	}
}

func ExecuteSelectQuery(db *pgx.Conn) {
	log.Println("starting execution of select query")
	//execute the query and get result rows
	rows, err := db.Query(context.Background(), "select * from public.person")
	if err != nil {
		log.Fatal("error while executing query")
	}

	log.Println("result:")
	//iterate through the rows
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}
		//convert DB types to Go types
		id := values[0].(int32)
		firstName := values[1].(string)
		lastName := values[2].(string)
		dateOfBirth := values[3].(time.Time)
		log.Println("[id:", id, ", first_name:", firstName, ", last_name:", lastName, ", date_of_birth:", dateOfBirth, "]")
	}

}

func ExecuteFunction(db *pgx.Conn) {
	log.Println("starting execution of databse function")
	// id can be taken as a user input
	// for the time being, let's hard code it
	id := 1

	//execute the query and get result rows
	rows, err := db.Query(context.Background(), "select * from public.get_person_details($1)", id)
	log.Println("input id: ", id)
	if err != nil {
		log.Fatal("error while executing query")
	}

	log.Println("result:")
	//iterate through the rows
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}

		//convert DB types to Go types
		firstName := values[0].(string)
		lastName := values[1].(string)
		dateOfBirth := values[2].(time.Time)

		log.Println("[first_name:", firstName, ", last_name:", lastName, ", date_of_birth:", dateOfBirth, "]")
	}

}
