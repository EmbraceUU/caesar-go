package pool

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

var (
	ErrInvalidPoolSize   = errors.New("invalid size for pool")
	ErrInvalidPoolExpiry = errors.New("invalid expiry for pool")
	ErrPoolClosed        = errors.New("pool is closed")
)

type sig struct{}

type f func() error

/*
思考：
1. 如何对worker进行管理
2. 如何定时清理过期worker
3. 如何保证并发时的读写不冲突
4. 分析每个结构体，为什么这么设计， 不能一味的用map和slice， 也不能一味的用chan
5. 运行中的goroutine如何通过release信号停止运行
*/
type Pool struct {
	// capacity of the pool
	capacity int32
	// the number of the currently running goroutines
	running int32
	// set the expired time (second) of every worker
	expiryDuration time.Duration
	// workers is a slice that store the available workers
	workers []*Worker
	// release is used to notice the pool to closed itself
	release chan sig
	// lock for synchronous operation
	lock sync.Mutex
	once sync.Once
}

func NewPool(size int) (*Pool, error) {
	return NewTimingPool(size, 10)
}

func NewTimingPool(size, expiry int) (*Pool, error) {
	if size <= 0 {
		return nil, ErrInvalidPoolSize
	}
	if expiry <= 0 {
		return nil, ErrInvalidPoolExpiry
	}

	p := &Pool{
		capacity:       int32(size),
		expiryDuration: time.Duration(expiry) * time.Second,
		release:        make(chan sig, 1),
	}
	return p, nil
}

func (p *Pool) Submit(task f) error {
	if len(p.release) > 0 {
		return ErrPoolClosed
	}
	w := p.getWorker()
	w.task <- task
	return nil
}

func (p *Pool) getWorker() *Worker {
	var w *Worker
	waiting := false
	p.lock.Lock()
	idleWorkers := p.workers
	n := len(idleWorkers) - 1

	if n < 0 {
		waiting = p.Running() >= p.Cap()
	} else {
		w = idleWorkers[n]
		idleWorkers[n] = nil
		p.workers = idleWorkers[:n]
	}
	p.lock.Unlock()

	if waiting {
		for {
			p.lock.Lock()
			idleWorkers = p.workers
			l := len(idleWorkers) - 1
			if l < 0 {
				p.lock.Unlock()
				continue
			}
			w = idleWorkers[l]
			idleWorkers[l] = nil
			p.workers = idleWorkers[:l]
			p.lock.Unlock()
			break
		}
	} else if w == nil {
		w := &Worker{
			pool: p,
			task: make(chan f, 1),
		}
		w.run()
		p.incRunning()
	}
	return w
}

func (p *Pool) putWorker(w *Worker) {
	w.recycleTime = time.Now()
	p.lock.Lock()
	p.workers = append(p.workers, w)
	p.lock.Unlock()
}

func (p *Pool) Running() int {
	return int(atomic.LoadInt32(&p.running))
}

func (p *Pool) Cap() int {
	return int(atomic.LoadInt32(&p.capacity))
}

func (p *Pool) incRunning() {
	atomic.AddInt32(&p.running, 1)
}

func (p *Pool) decRunning() {
	atomic.AddInt32(&p.running, -1)
}
