package handlers

import (
	"github.com/gabe565/portfolio/internal/handlers/github_readme_stats"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func RegisterLocalHandlers(e *core.ServeEvent, app core.App) error {
	e.Router.HEAD("/*", StaticHandler())
	e.Router.GET("/*", StaticHandler())

	e.Router.GET("/to/:handle", RedirectHandler(e), apis.ActivityLogger(app))

	if err := github_readme_stats.RegisterRoutes(e); err != nil {
		return err
	}

	return nil
}
