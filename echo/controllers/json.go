package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ReturnJSON(c echo.Context) error {
	data := new(interface{})

	if err := c.Bind(data); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, data)
}
