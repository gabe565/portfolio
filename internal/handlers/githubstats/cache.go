package githubstats

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strconv"
	"sync"
	"time"

	"github.com/pocketbase/pocketbase/core"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/svg"
)

func NewCache(sourceURL string) *Cache {
	return &Cache{sourceURL: sourceURL}
}

type Cache struct {
	sourceURL    string
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
	e.Response.Header().Set("Cache-Control", "public, max-age=86400, must-revalidate")
	e.Response.Header().Set("ETag", c.etag)
	http.ServeContent(e.Response, e.Request, "", c.lastModified, bytes.NewReader(c.data))
	return nil
}

func (c *Cache) RegisterRoutes(e *core.ServeEvent, endpoint string) *Cache {
	e.Router.HEAD(endpoint, c.Handler)
	e.Router.GET(endpoint, c.Handler)
	return c
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

	m := minify.New()
	m.AddFunc("image/svg+xml", svg.Minify)
	m.AddFunc("text/css", css.Minify)

	c.mu.RLock()
	buf := bytes.NewBuffer(make([]byte, 0, max(len(c.data), 2048)))
	c.mu.RUnlock()

	if err := m.Minify("image/svg+xml", buf, resp.Body); err != nil {
		return err
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = slices.Clip(buf.Bytes())
	c.lastModified = lastModified
	c.etag = etag
	return nil
}
