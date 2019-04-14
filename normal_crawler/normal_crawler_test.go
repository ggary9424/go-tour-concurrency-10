package normal_crawler

import (
	"os"
	"strconv"
	"testing"
)

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func TestNormalCrawler(t *testing.T) {
	depthStr := getenv("DEPTH", "10")

	depth, err := strconv.Atoi(depthStr)
	if err != nil {
		t.Fatal("Environment variable DEPTH should be integer.")
	}

	Crawl("https://golang.org/", depth, fetcher)
}
