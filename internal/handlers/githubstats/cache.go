package githubstats

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
)

func NewCache(endpoint, sourceURL string, interval time.Duration) *Cache {
	c := Cache{
		endpoint:  endpoint,
		sourceURL: sourceURL,
		interval:  interval,
	}
	go c.beginUpdate()
	return &c
}

type Cache struct {
	endpoint  string
	sourceURL string
	interval  time.Duration
	data      []byte
}

func (c *Cache) Handler(ctx echo.Context) error {
	if len(c.data) == 0 {
		return ctx.Redirect(http.StatusTemporaryRedirect, c.sourceURL)
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
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.sourceURL, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

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
