package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func Addfile(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	src, err := file.Open()
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	defer src.Close()
	Putfile(c.Request().Context(), src, file.Size, file.Filename)
	url := GetFileURL(c, file.Filename)
	fmt.Println(url)
	dst, err := os.Create(file.Filename)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.NoContent(http.StatusOK)
}

func GetFileURL(c echo.Context, file_name string) string {
	url, err := GeneratePresignedURL(c.Request().Context(), file_name)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return url
}
