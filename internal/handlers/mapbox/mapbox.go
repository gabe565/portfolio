package mapbox

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"gabe565.com/portfolio/internal/config"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func New(ctx context.Context, app *pocketbase.PocketBase, conf *config.Config) (*Client, error) {
	dir := filepath.Join(app.DataDir(), "maps")

	if err := os.MkdirAll(dir, 0o755); err != nil && !os.IsExist(err) {
		return nil, err
	}

	h := &Client{
		config: conf,
		app:    app,
		dir:    dir,
	}
	h.beginFetch(ctx)
	return h, nil
}

type Client struct {
	config *config.Config
	app    *pocketbase.PocketBase
	dir    string
}

type FetchRequest struct {
	Size  string
	Style string
	Zoom  float64
}

func (c *Client) beginFetch(ctx context.Context) {
	go func() {
		timer := time.NewTimer(0)
		defer timer.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-timer.C:
				timer.Reset(24 * time.Hour)
				if err := c.FetchAll(ctx); err != nil {
					slog.Error("Failed to download map", "error", err)
				}
			}
		}
	}()
}

func (c *Client) FetchAll(ctx context.Context) error {
	reqs := map[string]FetchRequest{
		"xs":  {"575x300", "gabe565/cm495kfdx00ip01rz7eifbthx", 9},
		"sm":  {"767x300", "gabe565/cm495kfdx00ip01rz7eifbthx", 9},
		"md":  {"991x300", "gabe565/cm495kfdx00ip01rz7eifbthx", 9},
		"lg":  {"1280x500", "gabe565/cm495kfdx00ip01rz7eifbthx", 9},
		"xl":  {"1280x500@2x", "gabe565/cm49gy50a014k01qr2sjoc52o", 9},
		"xxl": {"1280x500@2x", "gabe565/cm49gy50a014k01qr2sjoc52o", 8.5},
	}
	var errs []error
	var wg sync.WaitGroup
	var mu sync.Mutex
	for name, req := range reqs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := c.fetchMap(ctx, name, req); err != nil {
				mu.Lock()
				errs = append(errs, err)
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	return errors.Join(errs...)
}

func (c *Client) buildURL(conf FetchRequest) *url.URL {
	lat := strconv.FormatFloat(c.config.Map.Lat, 'f', -1, 64)
	lon := strconv.FormatFloat(c.config.Map.Lon, 'f', -1, 64)
	zoomStr := strconv.FormatFloat(conf.Zoom, 'f', -1, 64)

	u := &url.URL{
		Scheme: "https",
		Host:   "api.mapbox.com",
		Path:   path.Join("styles/v1", conf.Style, "static", lon+","+lat+","+zoomStr, conf.Size),
	}
	q := u.Query()
	q.Set("access_token", c.config.Map.Token)
	u.RawQuery = q.Encode()
	return u
}

var ErrFetchFailed = errors.New("failed to fetch map")

func (c *Client) fetchMap(ctx context.Context, name string, conf FetchRequest) error {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	u := c.buildURL(conf)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}
	if appURL := c.app.Settings().Meta.AppURL; appURL != "" {
		req.Header.Set("Referer", appURL)
	}

	etagPath := filepath.Join(c.dir, name+"_etag.txt")
	if etag, err := os.ReadFile(etagPath); err == nil {
		req.Header.Set("If-None-Match", string(etag))
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
		return fmt.Errorf("%w %q: %s", ErrFetchFailed, conf.Size, resp.Status)
	}

	tmp, err := os.Create(filepath.Join(c.dir, "."+name+".png"))
	if err != nil {
		return err
	}
	defer func() {
		_ = tmp.Close()
		_ = os.Remove(tmp.Name())
	}()

	if _, err := io.Copy(tmp, resp.Body); err != nil {
		return err
	}

	if err := tmp.Close(); err != nil {
		return err
	}

	if err := os.Rename(tmp.Name(), filepath.Join(c.dir, name+".png")); err != nil {
		return err
	}

	if etag := resp.Header.Get("ETag"); etag != "" {
		if err := os.WriteFile(etagPath, []byte(etag), 0o644); err != nil {
			return err
		}
	} else {
		_ = os.Remove(etagPath)
	}
	return nil
}

func (c *Client) Handler() func(e *core.RequestEvent) error {
	return apis.Static(os.DirFS(c.dir), false)
}
