package main

import "testing"

//BenchmarkNonBlock-8            1          4427172600 ns/op          6824 B/op             30 allocs/op
func BenchmarkNonBlock(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			NonBlock()
		}
	})
}

// BenchmarkBlock-8              56          20516550 ns/op          401619 B/op            2 allocs/op
func BenchmarkBlock(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Block()
		}
	})
}
