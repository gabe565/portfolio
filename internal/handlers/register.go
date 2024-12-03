package handlers

import (
	"context"

	"gabe565.com/portfolio/internal/config"
	"gabe565.com/portfolio/internal/handlers/githubstats"
	"github.com/pocketbase/pocketbase/core"
)

func RegisterLocalHandlers(ctx context.Context, conf *config.Config, e *core.ServeEvent) error {
	e.Router.GET("/{path...}", StaticHandler(conf))
	e.Router.GET("/to/{handle}", RedirectHandler())
	if err := githubstats.RegisterRoutes(ctx, conf, e); err != nil {
		return err
	}
	return nil
}
