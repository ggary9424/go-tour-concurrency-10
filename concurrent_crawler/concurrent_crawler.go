package concurrent_crawler

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type SafeCrawledUrlsRecorder struct {
	V   map[string]bool
	mux sync.Mutex
}

func (r *SafeCrawledUrlsRecorder) CheckIsExistedAndRecord(url string) bool {
	isExisted := false
	r.mux.Lock()
	if _, ok := r.V[url]; ok {
		isExisted = true
	} else {
		r.V[url] = true
	}
	r.mux.Unlock()
	return isExisted
}

func CrawlParallelly(url string, depth int, fetcher Fetcher, isDone chan bool, r *SafeCrawledUrlsRecorder) {
	if depth > 0 {
		if !r.CheckIsExistedAndRecord(url) {
			body, urls, err := fetcher.Fetch(url)
			if err != nil {
				fmt.Println(err)
				isDone <- true
				return
			}
			fmt.Printf("found: %s %q\n", url, body)

			innerIsDone := make(chan bool, len(urls))
			for _, u := range urls {
				go CrawlParallelly(u, depth, fetcher, innerIsDone, r)
			}

			for i := 0; i < len(urls); i++ {
				<-innerIsDone
			}
		}
	}
	isDone <- true
	return
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var F = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
