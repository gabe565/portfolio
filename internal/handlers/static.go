package handlers

import (
	"os"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals
var publicDir string

func Flags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&publicDir, "public", "frontend/dist", "Public directory")
}

func StaticHandler() echo.HandlerFunc {
	return apis.StaticDirectoryHandler(os.DirFS(publicDir), true)
}
