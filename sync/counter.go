package sync

import "sync"

type SyncCounter struct {
	mutex sync.Mutex
	v     int
}

func (c *SyncCounter) Inc() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.v++
}

func (c *SyncCounter) Value() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.v
}
