package handlers

import (
	"os"

	"gabe565.com/portfolio/internal/config"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func Static(conf *config.Config) func(*core.RequestEvent) error {
	return apis.Static(os.DirFS(conf.PublicDir), true)
}
