package githubstats

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
)

func NewCache(ctx context.Context, endpoint, sourceURL string, interval time.Duration) *Cache {
	c := Cache{
		endpoint:  endpoint,
		sourceURL: sourceURL,
		interval:  interval,
	}
	c.beginUpdate(ctx)
	return &c
}

type Cache struct {
	endpoint  string
	sourceURL string
	interval  time.Duration
	data      []byte
	mu        sync.RWMutex
}

func (c *Cache) Handler(ctx echo.Context) error {
	c.mu.RLock()
	defer c.mu.RUnlock()

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

var ErrUpstreamRequest = errors.New("upstream request failed")

func (c *Cache) Update(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
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
		return fmt.Errorf("%w: %s", ErrUpstreamRequest, resp.Status)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = b
	return nil
}

func (c *Cache) beginUpdate(ctx context.Context) {
	go func() {
		if err := c.Update(ctx); err != nil {
			log.Println(err)
		}

		ticker := time.NewTicker(c.interval)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				if err := c.Update(ctx); err != nil {
					log.Println(err)
				}
			}
		}
	}()
}
