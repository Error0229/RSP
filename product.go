package main

import (
	"fmt"
	"net/http"
	"rsp/simplesql"

	"github.com/labstack/echo/v4"
)

func AddProduct(c echo.Context) error {
	var product simplesql.AddProductParams
	err := c.Bind(&product)
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
