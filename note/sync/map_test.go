package sync

import (
	"testing"
)

func TestMapOp(t *testing.T) {
	MapOp()
}

func BenchmarkMapOpParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			MapOp()
		}
	})
}

func BenchmarkMapOpComParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			MapOpCom()
		}
	})
}
