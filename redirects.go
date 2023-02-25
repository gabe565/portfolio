package main

import (
	"database/sql"
	"errors"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
	"net/http"
)

func Redirects(e *core.ServeEvent) error {
	e.Router.GET("/to/:handle", func(c echo.Context) error {
		handle := c.PathParam("handle")

		record, err := e.App.Dao().FindFirstRecordByData("redirects", "handle", handle)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return c.JSON(http.StatusNotFound, map[string]any{
					"code":    http.StatusNotFound,
					"message": http.StatusText(http.StatusNotFound),
				})
			}
			return err
		}

		url := record.Get("url").(string)
		return c.Redirect(http.StatusTemporaryRedirect, url)
	})
	return nil
}
