package main

import (
	_ "github.com/gabe565/portfolio/migrations"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"log"
)

func main() {
	app := pocketbase.New()

	migratecmd.MustRegister(app, app.RootCmd, &migratecmd.Options{
		Automigrate: true,
	})

	app.OnBeforeServe().Add(FileHandler)
	app.OnBeforeServe().Add(Redirects)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
