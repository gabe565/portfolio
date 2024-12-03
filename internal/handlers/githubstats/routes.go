package githubstats

import (
	"context"
	"net/url"

	"gabe565.com/portfolio/internal/config"
	"github.com/pocketbase/pocketbase/core"
)

func RegisterRoutes(ctx context.Context, conf *config.Config, e *core.ServeEvent) error {
	parsedURL, err := url.Parse(conf.GitHubStats.SourceURL)
	if err != nil {
		return err
	}

	userURL, err := formatURL(parsedURL, "api", conf.GitHubStats.UserParams)
	if err != nil {
		return err
	}
	NewCache(ctx, "/api/github-stats/stats", userURL, conf.GitHubStats.Interval).RegisterRoutes(e)

	langsURL, err := formatURL(parsedURL, "api/top-langs", conf.GitHubStats.LangsParams)
	if err != nil {
		return err
	}
	NewCache(ctx, "/api/github-stats/top-langs", langsURL, conf.GitHubStats.Interval).RegisterRoutes(e)

	return nil
}

func formatURL(src *url.URL, path string, params map[string]string) (string, error) {
	u, err := src.Parse(path)
	if err != nil {
		return "", err
	}

	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	return u.String(), nil
}
