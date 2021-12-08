# go-postgresql-pgx-example
An example program to demonstrate PostgreSQL connectivity in Go using pgx library


before you run this, make sure you have -- 

1. go installed
2. postgres local host PostgreSQL 13.5
3. public db with a table called users as below

CREATE TABLE IF NOT EXISTS "users" (
	"id" VARCHAR NULL DEFAULT NULL,
	"name" VARCHAR NULL DEFAULT NULL,
	"email" VARCHAR NULL DEFAULT NULL,
	"age" INTEGER NULL DEFAULT NULL
);

INSERT INTO "users" ("id", "name", "email", "age") VALUES
	('1', 'keshava', 'keshava@gmail.com', 100),
	('2', 'madhava', 'madhava@gmail.com', 200);
