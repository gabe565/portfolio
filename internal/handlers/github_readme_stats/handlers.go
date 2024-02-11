package github_readme_stats

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

func CacheResponse(c echo.Context, data []byte, sourceUrl string) error {
	if len(data) == 0 {
		return c.Redirect(http.StatusTemporaryRedirect, sourceUrl)
	}

	c.Response().Header().Set("Content-Type", "image/svg+xml; charset=utf-8")
	c.Response().Header().Set("Cache-Control", "max-age="+strconv.Itoa(int(UpdateDuration.Seconds())))
	_, err := c.Response().Write(data)
	return err
}

func ReadmeStatsHandler(c echo.Context) error {
	return CacheResponse(c, ReadmeStatsCache, ReadmeStatsUrl)
}

func ReadmeTopLangsHandler(c echo.Context) error {
	return CacheResponse(c, TopLangsCache, TopLangsUrl)
}
