package pool

import (
	"sync"
	"time"
)

type sig struct {}

type f func() error

/*
思考：
1. 如何对worker进行管理
2. 如何定时清理过期worker
3. 如何保证并发时的读写不冲突
4. 分析每个结构体，为什么这么设计， 不能一味的用map和slice， 也不能一味的用chan
 */
type Pool struct {
	capacity int32
	running int32
	expiryDuration time.Duration
	workers []*Worker
	release chan sig
	lock sync.Mutex
	once sync.Once
}