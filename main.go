package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"rsp/simplesql"

	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, pass, dbname)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	queries := simplesql.New(db)
	test_user, err := queries.CreateUser(ctx, simplesql.CreateUserParams{
		Name: "kea", Email: "late@rsp.com", Password: "123", Role: "admin"})
	if err != nil {
		panic(err)
	}
	fmt.Println(test_user)

}
