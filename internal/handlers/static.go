package handlers

import (
	"os"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	flag "github.com/spf13/pflag"
)

var publicDir string

func init() {
	flag.StringVar(&publicDir, "public", "frontend/dist", "Public directory")
}

func StaticHandler() echo.HandlerFunc {
	return apis.StaticDirectoryHandler(os.DirFS(publicDir), true)
}
