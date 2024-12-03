package config

import "time"

type Config struct {
	PublicDir   string
	Turnstile   Turnstile
	GitHubStats GitHubStats
}

type Turnstile struct {
	Secret string
}

type GitHubStats struct {
	Interval    time.Duration
	SourceURL   string
	UserParams  map[string]string
	LangsParams map[string]string
}

func New() *Config {
	return &Config{
		PublicDir: "frontend/dist",
		Turnstile: Turnstile{
			Secret: "1x0000000000000000000000000000000AA",
		},
		GitHubStats: GitHubStats{
			Interval:  4 * time.Hour,
			SourceURL: "https://github-readme-stats.vercel.app",
			UserParams: map[string]string{
				"username":      "gabe565",
				"show_icons":    "true",
				"theme":         "transparent",
				"hide_border":   "true",
				"count_private": "true",
			},
			LangsParams: map[string]string{
				"username":      "gabe565",
				"show_icons":    "true",
				"theme":         "transparent",
				"hide_border":   "true",
				"count_private": "true",
			},
		},
	}
}
