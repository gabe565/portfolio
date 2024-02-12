package github_readme_stats

import (
	"github.com/pocketbase/pocketbase/core"
	flag "github.com/spf13/pflag"
)

func RegisterRoutes(e *core.ServeEvent) {
	updateInterval, err := flag.CommandLine.GetDuration(ReadmeStatsIntervalFlag)
	if err != nil {
		panic(err)
	}

	NewCache(
		"/api/github-stats/stats",
		"https://github-readme-stats.vercel.app/api?username=gabe565&show_icons=true&theme=transparent&hide_border=true&count_private=true",
		updateInterval,
	).RegisterRoutes(e)

	NewCache(
		"/api/github-stats/top-langs",
		"https://github-readme-stats.vercel.app/api/top-langs?username=gabe565&theme=transparent&hide_border=true&layout=compact",
		updateInterval,
	).RegisterRoutes(e)
}
