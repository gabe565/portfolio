package handlers

import (
	"os"

	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals
var publicDir string

func Flags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&publicDir, "public", "frontend/dist", "Public directory")
}

func StaticHandler() func(*core.RequestEvent) error {
	return apis.Static(os.DirFS(publicDir), true)
}
