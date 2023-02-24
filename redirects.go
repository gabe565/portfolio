package main

import (
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
	"net/http"
)

func Redirects(e *core.ServeEvent) error {
	e.Router.GET("/to/:handle", func(c echo.Context) error {
		handle := c.PathParam("handle")

		record, err := e.App.Dao().FindFirstRecordByData("redirects", "handle", handle)
		if err != nil {
			return err
		}

		url := record.Get("url").(string)
		return c.Redirect(http.StatusTemporaryRedirect, url)
	})
	return nil
}
