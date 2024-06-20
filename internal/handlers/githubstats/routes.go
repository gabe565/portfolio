package githubstats

import (
	"context"
	"net/url"

	"github.com/pocketbase/pocketbase/core"
)

func RegisterRoutes(ctx context.Context, e *core.ServeEvent) error {
	parsedURL, err := url.Parse(sourceURL)
	if err != nil {
		return err
	}

	userURL, err := formatURL(parsedURL, "api", userParams)
	if err != nil {
		return err
	}
	NewCache(ctx, "/api/github-stats/stats", userURL, interval).RegisterRoutes(e)

	langsURL, err := formatURL(parsedURL, "api/top-langs", langsParams)
	if err != nil {
		return err
	}
	NewCache(ctx, "/api/github-stats/top-langs", langsURL, interval).RegisterRoutes(e)

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
