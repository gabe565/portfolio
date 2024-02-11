package github_readme_stats

import (
	"io"
	"log"
	"net/http"
	"time"
)

var UpdateDuration = 4 * time.Hour

func init() {
	go func() {
		timer := time.NewTimer(0)
		for range timer.C {
			timer.Reset(UpdateDuration)
			if b, err := UpdateCache(ReadmeStatsUrl); err == nil {
				ReadmeStatsCache = b
			} else {
				log.Println(err)
			}
		}
	}()

	go func() {
		timer := time.NewTimer(0)
		for range timer.C {
			timer.Reset(UpdateDuration)
			if b, err := UpdateCache(TopLangsUrl); err == nil {
				TopLangsCache = b
			} else {
				log.Println(err)
			}
		}
	}()
}

var (
	ReadmeStatsUrl   = "https://github-readme-stats.vercel.app/api?username=gabe565&show_icons=true&theme=transparent&hide_border=true&count_private=true"
	ReadmeStatsCache []byte
)

var (
	TopLangsUrl   = "https://github-readme-stats.vercel.app/api/top-langs?username=gabe565&theme=transparent&hide_border=true&layout=compact"
	TopLangsCache []byte
)

func UpdateCache(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	return io.ReadAll(resp.Body)
}
