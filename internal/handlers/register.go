package handlers

import (
	"context"

	"gabe565.com/portfolio/internal/handlers/githubstats"
	"github.com/pocketbase/pocketbase/core"
)

func RegisterLocalHandlers(ctx context.Context, e *core.ServeEvent) error {
	e.Router.GET("/{path...}", StaticHandler())
	e.Router.GET("/to/{handle}", RedirectHandler())
	if err := githubstats.RegisterRoutes(ctx, e); err != nil {
		return err
	}
	return nil
}
