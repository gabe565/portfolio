package github_readme_stats

import (
	"bufio"
	"bytes"
	"io"
	"net/http"

	"github.com/labstack/echo/v5"
)

func CacheResponse(c echo.Context, cacheVal []byte, sourceUrl string) error {
	if ReadmeStatsCache == nil {
		return c.Redirect(http.StatusTemporaryRedirect, sourceUrl)
	}

	resp, err := http.ReadResponse(bufio.NewReader(bytes.NewReader(cacheVal)), nil)
	if err != nil {
		return err
	}

	c.Response().Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	c.Response().Header().Set("Cache-Control", resp.Header.Get("Cache-Control"))

	if _, err := io.Copy(c.Response(), resp.Body); err != nil {
		return err
	}
	return nil
}

func ReadmeStatsHandler(c echo.Context) error {
	return CacheResponse(c, ReadmeStatsCache, ReadmeStatsUrl)
}

func ReadmeTopLangsHandler(c echo.Context) error {
	return CacheResponse(c, TopLangsCache, TopLangsUrl)
}
