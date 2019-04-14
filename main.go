package main

import (
	crawler "github.com/ggary9424/go-tour-concurrency-10/concurrent_crawler"
)

func main() {
	r := &crawler.SafeCrawledUrlsRecorder{
		V: make(map[string]bool),
	}
	isDone := make(chan bool)
	go crawler.CrawlParallelly("https://golang.org/", 100, crawler.F, isDone, r)
	<-isDone
}
