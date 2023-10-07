package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
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

	e.POST("/product", AddProduct)

	e.POST("/file", Addfile)

	e.Logger.Fatal(e.Start(":8080"))
}
