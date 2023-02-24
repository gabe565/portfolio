package main

import (
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"os"
)

func FileHandler(e *core.ServeEvent) error {
	publicDir := os.Getenv("PUBLIC_DIR")
	if publicDir == "" {
		publicDir = "./frontend/dist"
	}

	e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS(publicDir), true))
	return nil
}
