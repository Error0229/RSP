package simplesql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var q *Queries

func GetQuerier() *Queries {
	if q == nil {
		user := os.Getenv("POSTGRES_USER")
		pass := os.Getenv("POSTGRES_PASSWORD")
		dbname := os.Getenv("POSTGRES_DB")
		dsn := fmt.Sprintf("host=127.0.0.1 user=%s password=%s dbname=%s sslmode=disable", user, pass, dbname)
		db, err := sql.Open("postgres", dsn)
		if err != nil {
			panic(err)
		}
		q = New(db)
	}

	return q
}
