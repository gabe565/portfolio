package main

import (
	"context"
	"log/slog"
	"os"

	"gabe565.com/portfolio/internal/captcha"
	"gabe565.com/portfolio/internal/config"
	"gabe565.com/portfolio/internal/contactform"
	"gabe565.com/portfolio/internal/handlers"
	"gabe565.com/portfolio/internal/handlers/githubstats"
	"gabe565.com/portfolio/internal/handlers/mapbox"
	_ "gabe565.com/portfolio/migrations"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {
	app := pocketbase.New()
	conf := config.New()
	conf.RegisterFlags(app.RootCmd)

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Automigrate: automigrateEnabled(),
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app.OnBootstrap().BindFunc(func(e *core.BootstrapEvent) error {
		if err := conf.Load(app.RootCmd); err != nil {
			return err
		}
		return e.Next()
	})

	app.OnServe().BindFunc(func(e *core.ServeEvent) error {
		slog.SetDefault(app.Logger())

		app.OnRecordCreateRequest("contact_form").BindFunc(captcha.Verify(conf))
		app.OnModelAfterCreateSuccess("contact_form").BindFunc(contactform.Notify(app))

		e.Router.BindFunc(func(e *core.RequestEvent) error {
			if e.Request.URL.Path == "/api/health" {
				return apis.SkipSuccessActivityLog().Func(e)
			}
			return e.Next()
		})

		e.Router.GET("/{path...}", handlers.Static(conf))
		e.Router.GET("/api/to/{handle}", handlers.Redirect())

		if err := githubstats.RegisterRoutes(ctx, conf, e); err != nil {
			return err
		}

		if conf.Map.Token != "" {
			mapClient, err := mapbox.New(ctx, app, conf)
			if err != nil {
				return err
			}

			e.Router.GET("/api/map/{path...}", mapClient.Handler())
		}

		return e.Next()
	})

	app.OnTerminate().BindFunc(func(e *core.TerminateEvent) error {
		cancel()
		return e.Next()
	})

	if err := app.Start(); err != nil {
		slog.Error("PocketBase returned an error", "error", err)
		os.Exit(1) //nolint:gocritic
	}
}
