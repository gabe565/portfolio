package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func RedirectHandler() func(*core.RequestEvent) error {
	return func(e *core.RequestEvent) error {
		handle := e.Request.PathValue("handle")

		record, err := e.App.FindFirstRecordByData("redirects", "handle", handle)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return apis.NewNotFoundError("", nil)
			}
			return err
		}

		url := record.Get("url").(string)
		return e.Redirect(http.StatusTemporaryRedirect, url)
	}
}
