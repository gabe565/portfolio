package config

import (
	"net/url"
	"time"
)

type Config struct {
	PublicDir   string
	Map         Map
	Turnstile   Turnstile
	GitHubStats GitHubStats
}

type Map struct {
	Token string
	Lat   float64
	Lon   float64
}

type Turnstile struct {
	Secret string
}

type GitHubStats struct {
	Interval  time.Duration
	SourceURL URL
	Username  string
}

func New() *Config {
	return &Config{
		PublicDir: "frontend/dist",
		Map: Map{
			Lat: 35.505,
			Lon: -97.4997,
		},
		Turnstile: Turnstile{
			Secret: "1x0000000000000000000000000000000AA",
		},
		GitHubStats: GitHubStats{
			Interval: 4 * time.Hour,
			SourceURL: URL{
				&url.URL{
					Scheme: "https",
					Host:   "github-readme-stats.vercel.app",
				},
			},
			Username: "gabe565",
		},
	}
}
