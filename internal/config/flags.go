package config

import (
	"github.com/spf13/cobra"
)

const (
	FlagPublic = "public"

	FlagMapToken = "mapbox-token"

	FlagTurnstileSecret = "turnstile-secret"

	FlagStatsInterval = "stats-interval"
	FlagStatsSource   = "stats-source"
	FlagStatsUsername = "stats-username"
)

func (c *Config) RegisterFlags(cmd *cobra.Command) {
	fs := cmd.PersistentFlags()
	fs.StringVar(&c.PublicDir, FlagPublic, c.PublicDir, "Public directory")

	fs.StringVar(&c.Map.Token, FlagMapToken, c.Map.Token, "Mapbox access token")

	fs.StringVar(&c.Turnstile.Secret, FlagTurnstileSecret, c.Turnstile.Secret, "Turnstile captcha secret key")

	fs.DurationVar(&c.GitHubStats.Interval, FlagStatsInterval, c.GitHubStats.Interval, "GitHub readme stats update interval")
	fs.Var(&c.GitHubStats.SourceURL, FlagStatsSource, "GitHub readme stats source URL")
	fs.StringVar(&c.GitHubStats.Username, FlagStatsUsername, c.GitHubStats.Username, "GitHub username to use for stats")
}
