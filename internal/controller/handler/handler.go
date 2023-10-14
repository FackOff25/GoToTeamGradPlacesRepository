package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateNotImplementedResponse(c echo.Context) error {
	defer c.Request().Body.Close()

	response := "Not implemented\n"

	return c.Blob(http.StatusNotImplemented, "plain/text", []byte(response))
}
