package githubstats

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"slices"
	"strconv"
	"sync"
	"time"

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
	endpoint     string
	sourceURL    string
	interval     time.Duration
	lastModified time.Time
	etag         string
	data         []byte
	mu           sync.RWMutex
}

func (c *Cache) Handler(e *core.RequestEvent) error {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if len(c.data) == 0 {
		return e.Redirect(http.StatusTemporaryRedirect, c.sourceURL)
	}

	e.Response.Header().Set("Content-Type", "image/svg+xml; charset=utf-8")
	e.Response.Header().Set("Cache-Control", "public, max-age="+strconv.Itoa(int(c.interval.Seconds()))+", must-revalidate")
	e.Response.Header().Set("ETag", c.etag)
	http.ServeContent(e.Response, e.Request, "", c.lastModified, bytes.NewReader(c.data))
	return nil
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

	if c.etag != "" {
		req.Header.Set("If-None-Match", c.etag)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_, _ = io.Copy(io.Discard, resp.Body)
		_ = resp.Body.Close()
	}()

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusNotModified:
		return nil
	default:
		return fmt.Errorf("%w: %s", ErrUpstreamRequest, resp.Status)
	}

	etag := resp.Header.Get("ETag")
	if etag != "" && etag == c.etag {
		return nil
	}

	lastModified := time.Now()
	if ageStr := resp.Header.Get("Age"); ageStr != "" {
		if age, err := strconv.Atoi(ageStr); err == nil {
			lastModified = lastModified.Add(-time.Duration(age) * time.Second)
		}
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	b = slices.Clip(b)

	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = b
	c.lastModified = lastModified
	c.etag = etag
	return nil
}

func (c *Cache) beginUpdate(ctx context.Context) {
	go func() {
		timer := time.NewTimer(0)
		defer timer.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-timer.C:
				timer.Reset(c.interval)
				if err := c.Update(ctx); err != nil {
					slog.Error("Failed to update GitHub stats", "error", err)
				}
			}
		}
	}()
}
