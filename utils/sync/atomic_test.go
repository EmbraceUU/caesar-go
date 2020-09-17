package sync

import "testing"

func TestCompareAndSwapInt32(t *testing.T) {
	a := int32(3)
	old := int32(0)
	n := int32(1)
	t.Log(CompareAndSwapInt32(&a, old, n))
	t.Log(a)
}
