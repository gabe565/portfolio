package handlers

import (
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"os"
)

func StaticHandler() echo.HandlerFunc {
	publicDir := os.Getenv("PUBLIC_DIR")
	if publicDir == "" {
		publicDir = "./frontend/dist"
	}

	return apis.StaticDirectoryHandler(os.DirFS(publicDir), true)
}
