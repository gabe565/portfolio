package main

import (
	"github.com/gabe565/portfolio/internal/contact_form"
	"github.com/gabe565/portfolio/internal/handlers"
	_ "github.com/gabe565/portfolio/migrations"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"log"
)

func main() {
	app := pocketbase.New()

	migratecmd.MustRegister(app, app.RootCmd, &migratecmd.Options{
		Automigrate: true,
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/*", handlers.StaticHandler())
		e.Router.GET("/to/:handle", handlers.RedirectHandler(e))
		return nil
	})

	app.OnModelAfterCreate("contact_form").Add(contact_form.Notify(app))

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
