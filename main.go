package main

import (
	"log"

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

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		return handlers.RegisterLocalHandlers(e, app)
	})

	app.OnRecordBeforeCreateRequest("contact_form").Add(captcha.Verify)
	app.OnModelAfterCreate("contact_form").Add(contactform.Notify(app))

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
