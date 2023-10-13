package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

// Addfile uploads a file to the server.
// @Summary Upload a file
// @Description Upload a file to the server
// @Tags Files
// @Accept mpfd
// @Produce json
// @Param file formData file true "File to upload"
// @Success 200 {string} string "OK"
// @Failure 400 {json} string "Bad Request"
// @Router /file [post]
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

// GetFile redirects the user to the URL of the specified file.
// @Summary Get a file
// @Description Redirect the user to the URL of the specified file
// @Tags Files
// @Param file_name path string true "Name of the file to retrieve"
// @Success 301 {string} string "Moved Permanently"
// @Failure 400 {string} string "Bad Request"
// @Router /file/{file_name} [get]
func GetFile(c echo.Context) error {
	fmt.Println("GetFile")
	file_name := c.Param("file_name")
	fmt.Println(file_name)
	url := GetFileURL(c, file_name)
	fmt.Println(url)
	return c.Redirect(http.StatusMovedPermanently, url)
}

func GetFileURL(c echo.Context, file_name string) string {
	url, err := GeneratePresignedURL(c.Request().Context(), file_name)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return url
}
