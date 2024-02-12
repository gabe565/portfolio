package github_readme_stats

import (
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
)

func NewCache(endpoint, sourceUrl string, interval time.Duration) *Cache {
	c := Cache{
		endpoint:  endpoint,
		sourceUrl: sourceUrl,
		interval:  interval,
	}
	go c.beginUpdate()
	return &c
}

type Cache struct {
	endpoint  string
	sourceUrl string
	interval  time.Duration
	data      []byte
}

func (c *Cache) Handler(ctx echo.Context) error {
	if len(c.data) == 0 {
		return ctx.Redirect(http.StatusTemporaryRedirect, c.sourceUrl)
	}

	ctx.Response().Header().Set("Content-Type", "image/svg+xml; charset=utf-8")
	ctx.Response().Header().Set("Cache-Control", "max-age="+strconv.Itoa(int(c.interval.Seconds())))
	_, err := ctx.Response().Write(c.data)
	return err
}

func (c *Cache) RegisterRoutes(e *core.ServeEvent) {
	e.Router.HEAD(c.endpoint, c.Handler)
	e.Router.GET(c.endpoint, c.Handler)
}

var ErrInvalidResponse = errors.New("invalid response")

func (c *Cache) Update() error {
	resp, err := http.Get(c.sourceUrl)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return ErrInvalidResponse
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	c.data = b
	return nil
}

func (c *Cache) beginUpdate() {
	if err := c.Update(); err != nil {
		log.Println(err)
	}

	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for range ticker.C {
		if err := c.Update(); err != nil {
			log.Println(err)
		}
	}
}
