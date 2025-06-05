package concurrency

import (
	"fmt"
	"time"
)

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go checkURLWithTimeout(wc, resultChannel, url)
	}

	for range len(urls) {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}

func checkURLWithTimeout(wc WebsiteChecker, resultChannel chan<- result, url string) {
	responseChan := make(chan bool, 1)

	go func() {
		responseChan <- wc(url)
	}()

	select {
	case valid := <-responseChan:
		resultChannel <- result{url, valid}
	case <-time.After(20 * time.Millisecond):
		fmt.Printf("Timeout for %s\n", url)
		resultChannel <- result{url, false}
	}
}
