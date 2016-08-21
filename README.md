# go-web-crawler
Solution to "Tour of Go" web crawler exercise using two different algorithms/approaches:

- sync : a recursive goroutine that uses a sycnhronized map to communicate
- channel : a function that launches multiple goroutines that communicate via channels

# Usage

    go build
  
Provide an -algorithm flag to choose which algorithm is used (optional - the default is "channel")

On Mac OSX/Linux

    ./go-web-crawler -algorithm=sync
    ./go-web-crawler -algorithm=channel
  
On Windows  

    ./go-web-crawle.exe -algorithm=sync
    ./go-web-crawler.exe -algorithm=channel

