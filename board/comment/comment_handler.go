package comment

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetComments(c echo.Context) error {
	return c.String(http.StatusOK, "GetComments")
}

func AddComment(c echo.Context) error {
	return c.String(http.StatusOK, "AddComment")
}

func EditComment(c echo.Context) error {
	return c.String(http.StatusOK, "EditComment")
}

func RemoveComment(c echo.Context) error {
	return c.String(http.StatusOK, "RemoveComment")
}
