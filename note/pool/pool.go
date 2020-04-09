package pool

import (
	"errors"
	"sync"
	"time"
)

var (
	ErrInvalidPoolSize   = errors.New("invalid size for pool")
	ErrInvalidPoolExpiry = errors.New("invalid expiry for pool")
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
