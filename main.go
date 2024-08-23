package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/gabe565/portfolio/internal/captcha"
	"github.com/gabe565/portfolio/internal/contactform"
	"github.com/gabe565/portfolio/internal/handlers"
	"github.com/gabe565/portfolio/internal/handlers/githubstats"
	_ "github.com/gabe565/portfolio/migrations"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {
	app := pocketbase.New()
	handlers.Flags(app.RootCmd)
	captcha.Flags(app.RootCmd)
	githubstats.Flags(app.RootCmd)

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Automigrate: automigrateEnabled(),
	})

	ctx, cancel := context.WithCancel(context.Background())

	app.OnTerminate().Add(func(_ *core.TerminateEvent) error {
		cancel()
		return nil
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		slog.SetDefault(app.Logger())
		return handlers.RegisterLocalHandlers(ctx, e, app)
	})

	app.OnRecordBeforeCreateRequest("contact_form").Add(captcha.Verify)
	app.OnModelAfterCreate("contact_form").Add(contactform.Notify(app))

	if err := app.Start(); err != nil {
		slog.Error("PocketBase returned an error", "error", err)
		os.Exit(1)
	}
}
