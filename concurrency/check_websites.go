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
		go func(urlInternal string) {
			responseChan := make(chan bool, 1)

			go func() {
				responseChan <- wc(urlInternal)
			}()

			select {
			case valid := <-responseChan:
				resultChannel <- result{urlInternal, valid}
			case <-time.After(20 * time.Millisecond):
				fmt.Printf("Timeout for %s\n", url)
				resultChannel <- result{urlInternal, false}
			}
		}(url)
	}

	for range len(urls) {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
