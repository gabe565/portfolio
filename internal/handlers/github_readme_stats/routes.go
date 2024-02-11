package github_readme_stats

import "github.com/pocketbase/pocketbase/core"

func RegisterRoutes(e *core.ServeEvent) {
	go beginUpdater()

	e.Router.HEAD("/api/github-stats/stats", ReadmeStatsHandler)
	e.Router.GET("/api/github-stats/stats", ReadmeStatsHandler)

	e.Router.HEAD("/api/github-stats/top-langs", ReadmeTopLangsHandler)
	e.Router.GET("/api/github-stats/top-langs", ReadmeTopLangsHandler)
}
