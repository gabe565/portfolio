package captcha

import (
	"errors"

	"gabe565.com/portfolio/internal/config"
	"github.com/meyskens/go-turnstile"
	"github.com/pocketbase/pocketbase/core"
)

var ErrInvalidCaptcha = errors.New("invalid captcha")

func Verify(conf *config.Config) func(*core.RecordRequestEvent) error {
	return func(e *core.RecordRequestEvent) error {
		if conf.Turnstile.Secret != "" {
			ts := turnstile.New(conf.Turnstile.Secret)
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
}
