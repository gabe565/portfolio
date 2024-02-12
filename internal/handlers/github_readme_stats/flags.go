package github_readme_stats

import (
	"time"

	flag "github.com/spf13/pflag"
)

var (
	interval    time.Duration
	sourceUrl   string
	userParams  = make(map[string]string)
	langsParams = make(map[string]string)
)

func init() {
	flag.DurationVar(&interval, "stats-interval", 4*time.Hour, "GitHub readme stats update interval")
	flag.StringVar(&sourceUrl, "stats-source", "https://github-readme-stats.vercel.app", "GitHub readme stats source URL")

	flag.StringToStringVar(&userParams, "stats-user-params", map[string]string{
		"username":      "gabe565",
		"show_icons":    "true",
		"theme":         "transparent",
		"hide_border":   "true",
		"count_private": "true",
	}, "GitHub readme stats params")

	flag.StringToStringVar(&langsParams, "stats-langs-params", map[string]string{
		"username":    "gabe565",
		"theme":       "transparent",
		"hide_border": "true",
		"layout":      "compact",
	}, "GitHub readme stats top-langs params")
}
