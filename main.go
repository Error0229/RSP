package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"rsp/simplesql"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// export POSTGRESQL_URL='postgres://root:1234@localhost:5432/db?sslmode=disable'
// migrate create -ext sql -dir db/migrations -seq create_product_table
// migrate -path db/migrations/ -database $POSTGRESQL_URL up
func main() {

	e := echo.New()
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	dsn := fmt.Sprintf("host=127.0.0.1 user=%s password=%s dbname=%s sslmode=disable", user, pass, dbname)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	queries := simplesql.New(db)

	e.POST("/product", func(c echo.Context) error {
		var product simplesql.AddProductParams
		err := c.Bind(&product)
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusBadRequest, err)
		}
		_, err = queries.AddProduct(c.Request().Context(), product)
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusBadRequest, err)
		}

		fmt.Printf("%+v\n", product)

		return c.NoContent(http.StatusOK)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
