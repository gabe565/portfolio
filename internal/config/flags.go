package config

import (
	"github.com/spf13/cobra"
)

const (
	FlagPublic = "public"

	FlagMapToken = "mapbox-token"

	FlagTurnstileSecret = "turnstile-secret"

	FlagStatsInterval    = "stats-interval"
	FlagStatsSource      = "stats-source"
	FlagStatsUserParams  = "stats-user-params"
	FlagStatsLangsParams = "stats-langs-params"
)

func (c *Config) RegisterFlags(cmd *cobra.Command) {
	fs := cmd.PersistentFlags()
	fs.StringVar(&c.PublicDir, FlagPublic, c.PublicDir, "Public directory")

	fs.StringVar(&c.Map.Token, FlagMapToken, c.Map.Token, "Mapbox access token")

	fs.StringVar(&c.Turnstile.Secret, FlagTurnstileSecret, c.Turnstile.Secret, "Turnstile captcha secret key")

	fs.DurationVar(&c.GitHubStats.Interval, FlagStatsInterval, c.GitHubStats.Interval, "GitHub readme stats update interval")
	fs.Var(&c.GitHubStats.SourceURL, FlagStatsSource, "GitHub readme stats source URL")
	fs.StringToStringVar(&c.GitHubStats.UserParams, FlagStatsUserParams, c.GitHubStats.UserParams, "GitHub readme stats params")
	fs.StringToStringVar(&c.GitHubStats.LangsParams, FlagStatsLangsParams, c.GitHubStats.LangsParams, "GitHub readme stats top-langs params")
}
