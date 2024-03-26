package githubstats

import (
	"time"

	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals
var (
	interval    time.Duration
	sourceURL   string
	userParams  = make(map[string]string)
	langsParams = make(map[string]string)
)

func Flags(cmd *cobra.Command) {
	cmd.PersistentFlags().DurationVar(&interval, "stats-interval", 4*time.Hour, "GitHub readme stats update interval")
	cmd.PersistentFlags().StringVar(&sourceURL, "stats-source", "https://github-readme-stats.vercel.app", "GitHub readme stats source URL")

	cmd.PersistentFlags().StringToStringVar(&userParams, "stats-user-params", map[string]string{
		"username":      "gabe565",
		"show_icons":    "true",
		"theme":         "transparent",
		"hide_border":   "true",
		"count_private": "true",
	}, "GitHub readme stats params")

	cmd.PersistentFlags().StringToStringVar(&langsParams, "stats-langs-params", map[string]string{
		"username":    "gabe565",
		"theme":       "transparent",
		"hide_border": "true",
		"layout":      "compact",
	}, "GitHub readme stats top-langs params")
}
