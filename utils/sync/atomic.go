package sync

import "sync/atomic"

func CompareAndSwapInt32(data *int32, old, new int32) bool {
	// 如果data==old, 返回true
	// 如果data==old, data=new
	return atomic.CompareAndSwapInt32(data, old, new)
}
