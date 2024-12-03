package handlers

import (
	"context"

	"gabe565.com/portfolio/internal/config"
	"gabe565.com/portfolio/internal/handlers/githubstats"
	"github.com/pocketbase/pocketbase/core"
)

func Register(ctx context.Context, conf *config.Config, e *core.ServeEvent) error {
	e.Router.GET("/{path...}", Static(conf))
	e.Router.GET("/to/{handle}", Redirect())
	if err := githubstats.RegisterRoutes(ctx, conf, e); err != nil {
		return err
	}
	return nil
}
