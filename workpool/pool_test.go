package workpool

import (
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := New(5)

	for i := range 10 {
		err := pool.Schedule(func() {
			time.Sleep(time.Second * 3)
		})
		if err != nil {
			println("task: ", i, " error: ", err)
		}
	}

	pool.Free()
}

func TestPoolWithOptions(t *testing.T) {
	pool := New(5, WithPreAllocWorkers(true), WithBlock(true))

	for i := range 10 {
		err := pool.Schedule(func() {
			time.Sleep(time.Second * 3)
		})
		if err != nil {
			println("task: ", i, " error: ", err)
		}
	}

	pool.Free()
}
