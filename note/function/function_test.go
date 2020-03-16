package function

import "testing"

func TestDeferPractice(t *testing.T) {
	DeferPractice(0)
}

func TestDeferPractice1(t *testing.T) {
	DeferPractice1()
}

/*
BenchmarkDeferPractice2-4   	30000000	        51.9 ns/op
BenchmarkDeferPractice2-4   	100000000	        15.8 ns/op
*/
func BenchmarkDeferPractice2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DeferPractice2()
	}
}
