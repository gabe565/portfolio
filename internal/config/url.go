package config

import "net/url"

type URL struct {
	*url.URL
}

func (u *URL) Set(s string) error {
	newURL, err := url.Parse(s)
	if err != nil {
		return err
	}
	u.URL = newURL
	return err
}

func (u *URL) Type() string {
	return "string"
}
