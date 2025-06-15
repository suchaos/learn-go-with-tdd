package sync

import (
	"sync"
	"testing"
)

func TestSyncCounter(t *testing.T) {
	counter := &SyncCounter{}

	var wg sync.WaitGroup

	for range 10 {
		wg.Add(1)

		go func() {
			for range 100000 {
				counter.Inc()
			}
			wg.Done()
		}()
	}

	wg.Wait()

	got := counter.Value()
	want := 100000 * 10

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
