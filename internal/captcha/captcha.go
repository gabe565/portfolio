package captcha

import (
	"errors"
	"os"

	"github.com/meyskens/go-turnstile"
	"github.com/pocketbase/pocketbase/core"
	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals
var turnstileSecret string

func Flags(cmd *cobra.Command) {
	defaultTurnstileSecret := "1x0000000000000000000000000000000AA"
	if env := os.Getenv("TURNSTILE_SECRET"); env != "" {
		defaultTurnstileSecret = env
	}

	cmd.PersistentFlags().StringVar(
		&turnstileSecret,
		"turnstile-secret",
		defaultTurnstileSecret,
		"Turnstile captcha secret key",
	)
}

var ErrInvalidCaptcha = errors.New("invalid captcha")

func Verify(e *core.RecordRequestEvent) error {
	if turnstileSecret != "" {
		ts := turnstile.New(turnstileSecret)
		val := e.Request.Header.Get("X-Captcha")

		resp, err := ts.Verify(val, e.RealIP())
		if err != nil {
			return err
		}

		if !resp.Success {
			return ErrInvalidCaptcha
		}
	}
	return e.Next()
}
