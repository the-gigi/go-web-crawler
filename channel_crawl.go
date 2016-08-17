package main

// var visited := make(map[string]bool)

func fetchURL(url string, ch chan int) {
	// 	body, urls, err := fetcher.Fetch(url)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
}

// ChannelCrawl uses fetcher to crawl
// pages starting with url, to a maximum of depth.
func ChannelCrawl(url string, depth int, fetcher Fetcher) {
	// 	defer wg.Done()

	// 	if depth <= 0 {
	// 		return
	// 	}

	// 	// Check if this url has already been fetched (or being fetched)
	// 	fetchedUrls.m.Lock()
	// 	if fetchedUrls.urls[url] {
	// 		fetchedUrls.m.Unlock()
	// 		return
	// 	}

	// 	// OK. eEt's fetch this url
	// 	fetchedUrls.urls[url] = true
	// 	fetchedUrls.m.Unlock()

	// 	body, urls, err := fetcher.Fetch(url)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	fmt.Printf("found: %s %q\n", url, body)
	// 	for _, u := range urls {
	// 		wg.Add(1)
	// 		go Crawl(u, depth-1, fetcher)
	// 	}
	// 	return
}
