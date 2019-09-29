package hmath

import (
	"math/rand"
	"testing"
)

// 单元测试
// 命名: TestFunc
// 入参: testing.T
func TestCombination(t *testing.T) {
	data := []struct {
		m        int
		n        int
		expected int
	}{
		{1, 0, 1},
		{4, 1, 4},
		{4, 2, 6},
		{4, 3, 4},
		{4, 4, 1},
		{10, 1, 10},
		{10, 3, 120},
		{10, 7, 120},
	}

	for _, unit := range data {
		if actually := Combination(unit.m, unit.n); actually != unit.expected {
			t.Errorf("combination: [%v], actually: [%v]", unit, actually)
		}
	}
}

// 性能测试
// 命名: BenchmarkFunc
// 入参: testing.B
func BenchmarkCombination(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Combination(i+1, rand.Intn(i+1))
	}
}

// 并发性能测试
// 命名: BenchmarkFuncParallel
// 入参: testing.B
func BenchmarkCombinationParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m := rand.Intn(100) + 1
			n := rand.Intn(m)
			Combination(m, n)
		}
	})
}
