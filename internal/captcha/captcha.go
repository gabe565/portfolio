package captcha

import (
	"errors"
	"os"

	"github.com/meyskens/go-turnstile"
	"github.com/pocketbase/pocketbase/core"
	flag "github.com/spf13/pflag"
)

var turnstileSecret string

func init() {
	flag.StringVar(
		&turnstileSecret,
		"turnstile-secret",
		os.Getenv("TURNSTILE_SECRET"),
		"Turnstile captcha secret key",
	)
}

var ErrInvalidCaptcha = errors.New("invalid captcha")

func Verify(e *core.RecordCreateEvent) error {
	if turnstileSecret != "" {
		ts := turnstile.New(turnstileSecret)
		val := e.HttpContext.Request().Header.Get("X-Captcha")

		resp, err := ts.Verify(val, e.HttpContext.RealIP())
		if err != nil {
			return err
		}

		if !resp.Success {
			return ErrInvalidCaptcha
		}
	}
	return nil
}
