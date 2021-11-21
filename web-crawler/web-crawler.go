package main

import (
	"fmt"
	"sync"
)

type SeenUrls struct {
	mu   sync.Mutex
	urls map[string]int
}

func (s *SeenUrls) Inc(url string) {
	s.mu.Lock()
	s.urls[url]++
	s.mu.Unlock()
}

func (s *SeenUrls) Seen(url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.urls[url]
	return ok
}

var seenUrls = SeenUrls{urls: make(map[string]int)}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, ch chan string) {
	defer close(ch)
	if seenUrls.Seen(url) {
		return
	} else {
		seenUrls.Inc(url)
	}

	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	ch <- fmt.Sprintf("found: %s %q\n", url, body)

	result := make([]chan string, len(urls))

	for i, u := range urls {
		result[i] = make(chan string)
		go Crawl(u, depth-1, fetcher, result[i])
	}
	for i := range result {
		for s := range result[i] {
			ch <- s
		}
	}

	return
}

func main() {
	ch := make(chan string)
	go Crawl("https://golang.org/", 4, fetcher, ch)
	for s := range ch {
		fmt.Println(s)
	}
	//fmt.Println(<-ch)
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
