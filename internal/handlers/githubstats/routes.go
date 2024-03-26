package githubstats

import (
	"net/url"

	"github.com/pocketbase/pocketbase/core"
)

func RegisterRoutes(e *core.ServeEvent) error {
	parsedURL, err := url.Parse(sourceURL)
	if err != nil {
		return err
	}

	userURL, err := formatURL(parsedURL, "api", userParams)
	if err != nil {
		return err
	}
	NewCache("/api/github-stats/stats", userURL, interval).RegisterRoutes(e)

	langsURL, err := formatURL(parsedURL, "api/top-langs", langsParams)
	if err != nil {
		return err
	}
	NewCache("/api/github-stats/top-langs", langsURL, interval).RegisterRoutes(e)

	return nil
}

func formatURL(src *url.URL, parse string, params map[string]string) (string, error) {
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
