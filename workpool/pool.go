package workpool

import (
	"errors"
	"fmt"
	"sync"
)

var (
	ErrNoIdleWorkerInPool = errors.New("no idle worker in pool") // workerpool中任务已满，没有空闲goroutine用于处理新任务
	ErrWorkerPoolFreed    = errors.New("workerpool freed")       // workerpool已终止运行
)

type Pool struct {
	capacity int
	// 是否在创建pool的时候就预创建workers，默认值为：false
	preAlloc bool

	// 当pool满的情况下，新的Schedule调用是否阻塞当前goroutine。默认值：true
	// 如果block = false，则Schedule返回ErrNoWorkerAvailInPool
	block  bool
	active chan struct{}

	tasks chan Task

	wg   sync.WaitGroup
	quit chan struct{}
}

type Task func()

const (
	defaultCapacity = 100
	maxCapacity     = 10000
)

func New(capacity int, opts ...Option) *Pool {
	if capacity <= 0 {
		capacity = defaultCapacity
	}
	if capacity > maxCapacity {
		capacity = maxCapacity
	}

	p := &Pool{
		capacity: capacity,
		preAlloc: false,
		tasks:    make(chan Task),
		block:    true,
		quit:     make(chan struct{}),
		active:   make(chan struct{}, capacity),
	}

	for _, opt := range opts {
		opt(p)
	}

	fmt.Printf("workpool start(preAlloc=%t)\n", p.preAlloc)
	if p.preAlloc {
		for i := 0; i < p.capacity; i++ {
			p.newWorker(i)
			p.active <- struct{}{}
		}
	}

	go p.run()

	return p
}

func (p *Pool) run() {
	idx := len(p.active)

	if !p.preAlloc {
	loop:
		for task := range p.tasks {
			p.returnTask(task)
			select {
			case <-p.quit:
				return
			case p.active <- struct{}{}:
				idx++
				p.newWorker(idx)
			default:
				break loop
			}
		}
	}

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

func (p *Pool) Schedule(task Task) error {
	select {
	case <-p.quit:
		return ErrWorkerPoolFreed
	case p.tasks <- task:
		return nil
	default:
		if p.block {
			p.tasks <- task
			return nil
		}
		return ErrNoIdleWorkerInPool
	}
}

func (p *Pool) Free() {
	close(p.quit)
	p.wg.Wait()
	fmt.Printf("worker pool freed\n")
}

func (p *Pool) returnTask(task Task) {
	go func() {
		p.tasks <- task
	}()
}
