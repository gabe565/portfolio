package handlers

import (
	"context"

	"github.com/gabe565/portfolio/internal/handlers/githubstats"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func RegisterLocalHandlers(ctx context.Context, e *core.ServeEvent, app core.App) error {
	static := StaticHandler()
	e.Router.HEAD("/*", static)
	e.Router.GET("/*", static)

	e.Router.GET("/to/:handle", RedirectHandler(e), apis.ActivityLogger(app))

	if err := githubstats.RegisterRoutes(ctx, e); err != nil {
		return err
	}

	return nil
}
