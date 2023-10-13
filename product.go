package main

import (
	"fmt"
	"net/http"
	"rsp/simplesql"

	"github.com/labstack/echo/v4"
)

// AddProduct adds a new product to the database.
// @Summary Add a new product
// @Description Add a new product to the database
// @Tags Products
// @Accept json
// @Produce json
// @Param product body simplesql.AddProductParams true "Product details"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Router /product [post]
func AddProduct(c echo.Context) error {
	var product simplesql.AddProductParams
	err := c.Bind(&product)
	fmt.Println(product)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	_, err = simplesql.GetQuerier().AddProduct(c.Request().Context(), product)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	fmt.Printf("%+v\n", product)

	return c.NoContent(http.StatusOK)
}
