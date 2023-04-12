package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func RedirectHandler(e *core.ServeEvent) echo.HandlerFunc {
	return func(c echo.Context) error {
		handle := c.PathParam("handle")

		record, err := e.App.Dao().FindFirstRecordByData("redirects", "handle", handle)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return apis.NewNotFoundError("", nil)
			}
			return err
		}

		url := record.Get("url").(string)
		return c.Redirect(http.StatusTemporaryRedirect, url)
	}
}
