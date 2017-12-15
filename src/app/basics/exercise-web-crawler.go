package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	crawledUrls := make(map[string]int)
	var mux sync.Mutex
	ch := make(chan bool, 10)

	var crawler func(string, int, Fetcher)
	crawler = func(urlC string, depthC int, fetcherC Fetcher) {
		ch <- true
		defer func() {
			<-ch
		}()

		if depthC <= 0 {
			return
		}

		mux.Lock(); defer mux.Unlock()
		_, exists := crawledUrls[urlC]
		if exists {
			return
		}
		crawledUrls[urlC] = 1

		body, urls, err := fetcher.Fetch(urlC)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q\n", urlC, body)
		for _, u := range urls {
			go crawler(u, depth-1, fetcher)
		}
		return
	}
	go crawler(url, depth, fetcher)

	for {
		time.Sleep(100 * time.Millisecond)
		if 0 >= len(ch) {
			break
		}
	}
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
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
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
