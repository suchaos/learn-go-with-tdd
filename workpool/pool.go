package workpool

import (
	"errors"
	"fmt"
	"sync"
)

type Pool struct {
	capacity int

	active chan struct{}
	tasks  chan Task

	wg   sync.WaitGroup
	quit chan struct{}
}

type Task func()

const (
	defaultCapacity = 100
	maxCapacity     = 10000
)

func New(capacity int) *Pool {
	if capacity <= 0 {
		capacity = defaultCapacity
	}
	if capacity > maxCapacity {
		capacity = maxCapacity
	}

	p := &Pool{
		capacity: capacity,
		tasks:    make(chan Task),
		quit:     make(chan struct{}),
		active:   make(chan struct{}, capacity),
	}

	fmt.Println("workpool.New: capacity:", capacity)

	go p.run()

	return p
}

func (p *Pool) run() {
	idx := 0
	for {
		select {
		case <-p.quit:
			return
		case p.active <- struct{}{}:
			// active 可写时, 创建一个新的worker
			idx++
			p.newWorker(idx)
		}
	}
}

func (p *Pool) newWorker(i int) {
	p.wg.Add(1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("worker[%03d]: recover panic[%s] and exit", i, err)
				<-p.active
			}
			p.wg.Done()
		}()

		fmt.Printf("worker[%03d]: start\n", i)

		for {
			select {
			case <-p.quit:
				fmt.Printf("worker[%03d]: quit\n", i)
				<-p.active
				return
			case task := <-p.tasks:
				fmt.Printf("worker[%03d]: receive a task\n", i)
				task()
			}
		}
	}()
}

var ErrWorkerPoolFreed = errors.New("worker pool freed")

func (p *Pool) Schedule(task Task) error {
	select {
	case <-p.quit:
		return ErrWorkerPoolFreed
	case p.tasks <- task:
		return nil
	}
}

func (p *Pool) Free() {
	close(p.quit)
	p.wg.Wait()
	fmt.Printf("worker pool freed\n")
}
