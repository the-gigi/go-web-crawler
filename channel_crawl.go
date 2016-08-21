package main

import "fmt"

type links struct {
	urls  []string
	depth int
}

func fetchURL(url string, depth int, candidates chan links) {
	body, urls, err := fetcher.Fetch(url)
	fmt.Printf("found: %s %q\n", url, body)

	if err != nil {
		fmt.Println(err)
	}

	candidates <- links{urls, depth - 1}
}

// ChannelCrawl crawls links from a seed url
func ChannelCrawl(url string, depth int, fetcher Fetcher) {
	candidates := make(chan links, 0)
	fetched := make(map[string]bool)
	counter := 1

	// Fetch initial url to seed the candidates channel
	go fetchURL(url, depth, candidates)

	for counter > 0 {
		candidateLinks := <-candidates
		counter--
		depth = candidateLinks.depth
		for _, candidate := range candidateLinks.urls {
			// Already fetched. Continue...
			if fetched[candidate] {
				continue
			}

			// Add to fetched mapped
			fetched[candidate] = true

			if depth > 0 {
				counter++
				go fetchURL(candidate, depth, candidates)
			}
		}
	}
}
