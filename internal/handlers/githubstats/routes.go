package githubstats

import (
	"context"
	"log/slog"
	"net/url"

	"gabe565.com/portfolio/internal/config"
	"github.com/pocketbase/pocketbase/core"
)

func RegisterRoutes(ctx context.Context, conf *config.Config, e *core.ServeEvent) error {
	stats := NewCache(
		formatURL(conf.GitHubStats.SourceURL.URL, "api", conf.GitHubStats.UserParams),
	).RegisterRoutes(e, "/api/github-stats/stats")

	topLangs := NewCache(
		formatURL(conf.GitHubStats.SourceURL.URL, "api/top-langs", conf.GitHubStats.LangsParams),
	).RegisterRoutes(e, "/api/github-stats/top-langs")

	update := func() {
		go func() {
			if err := stats.Update(ctx); err != nil {
				slog.Error("Failed to update GitHub stats", "error", err)
			}
		}()
		go func() {
			if err := topLangs.Update(ctx); err != nil {
				slog.Error("Failed to update GitHub top langs", "error", err)
			}
		}()
	}

	go update()
	return e.App.Cron().Add("updateGitHubStats", "0 */4 * * *", update)
}

func formatURL(src *url.URL, path string, params map[string]string) string {
	u, err := src.Parse(path)
	if err != nil {
		panic(err)
	}

	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	return u.String()
}
