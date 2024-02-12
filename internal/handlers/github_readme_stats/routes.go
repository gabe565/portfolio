package github_readme_stats

import (
	"net/url"

	"github.com/pocketbase/pocketbase/core"
)

func RegisterRoutes(e *core.ServeEvent) error {
	parsedUrl, err := url.Parse(sourceUrl)
	if err != nil {
		return err
	}

	userUrl, err := formatUrl(parsedUrl, "api", userParams)
	if err != nil {
		return err
	}
	NewCache("/api/github-stats/stats", userUrl, interval).RegisterRoutes(e)

	langsUrl, err := formatUrl(parsedUrl, "api/top-langs", langsParams)
	if err != nil {
		return err
	}
	NewCache("/api/github-stats/top-langs", langsUrl, interval).RegisterRoutes(e)

	return nil
}

func formatUrl(src *url.URL, parse string, params map[string]string) (string, error) {
	parsed, err := src.Parse(parse)
	if err != nil {
		return parsed.String(), err
	}

	q := parsed.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	parsed.RawQuery = q.Encode()

	return parsed.String(), nil
}
