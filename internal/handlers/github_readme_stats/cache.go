package github_readme_stats

import (
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

func init() {
	go func() {
		if err := UpdateReadmeStatsCache(); err != nil {
			log.Println(err)
		}
		ticker := time.NewTicker(4 * time.Hour)
		for range ticker.C {
			if err := UpdateReadmeStatsCache(); err != nil {
				log.Println(err)
			}
		}
	}()

	go func() {
		if err := UpdateTopLangsCache(); err != nil {
			log.Println(err)
		}
		ticker := time.NewTicker(4 * time.Hour)
		for range ticker.C {
			if err := UpdateTopLangsCache(); err != nil {
				log.Println(err)
			}
		}
	}()
}

type CacheEntry struct {
	Headers http.Header
}

var (
	ReadmeStatsUrl   = "https://github-readme-stats.vercel.app/api?username=gabe565&show_icons=true&theme=transparent&hide_border=true&count_private=true"
	ReadmeStatsCache []byte
)

var (
	TopLangsUrl   = "https://github-readme-stats.vercel.app/api/top-langs?username=gabe565&theme=transparent&hide_border=true&layout=compact"
	TopLangsCache []byte
)

func UpdateReadmeStatsCache() error {
	resp, err := http.Get(ReadmeStatsUrl)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatal(err)
	}

	ReadmeStatsCache = dump

	return nil
}

func UpdateTopLangsCache() error {
	resp, err := http.Get(TopLangsUrl)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatal(err)
	}

	TopLangsCache = dump

	return nil
}
