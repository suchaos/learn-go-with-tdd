package concurrency

import (
	"maps"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	//time.Sleep(30 * time.Millisecond)
	return url != "always return ture"
}

func TestCheckWebsites(t *testing.T) {
	urls := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"http://amazon.com",
	}
	got := CheckWebsites(mockWebsiteChecker, urls)

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"http://amazon.com":          true,
	}

	if !maps.Equal(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := range 100 {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
