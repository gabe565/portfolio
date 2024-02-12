package github_readme_stats

import (
	"time"

	flag "github.com/spf13/pflag"
)

const (
	ReadmeStatsIntervalFlag = "readme-stats-interval"
)

func init() {
	flag.Duration(ReadmeStatsIntervalFlag, 4*time.Hour, "GitHub readme stats update interval")
}
