package main

import (
	"fmt"
	"sync"
)

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func crawl(url string, depth int, fetcher Fetcher) {
	defer wg.Done()

	if depth <= 0 {
		return
	}

	// Check if this url has already been fetched (or being fetched)
	fetchedUrls.m.Lock()
	if fetchedUrls.urls[url] {
		fetchedUrls.m.Unlock()
		return
	}

	// OK. eEt's fetch this url
	fetchedUrls.urls[url] = true
	fetchedUrls.m.Unlock()

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		wg.Add(1)
		go crawl(u, depth-1, fetcher)
	}
	return
}

var fetchedUrls = struct {
	urls map[string]bool
	m    sync.Mutex
}{urls: make(map[string]bool)}

// SyncMapCrawl uses a recursive crawl() functon to fetch all urls
func SyncMapCrawl(url string, depth int, fetcher Fetcher) {
	wg.Add(1)
	go crawl("http://golang.org/", 4, fetcher)
	wg.Wait()
}
