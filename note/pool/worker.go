package pool

import "time"

type Worker struct {
	// pool who owns this worker
	// 这里包含了pool的指针，而pool里面也有worker的指针
	pool *Pool
	task chan f
	recycleTime time.Time
}