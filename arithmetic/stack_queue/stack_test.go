package stack_queue

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	s := Constructor()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	s.GetMin()
	s.Pop()
	s.Top()
	s.GetMin()
}

func TestDecodeString(t *testing.T) {
	s := "3[a]2[bc]"
	r := DecodeString(s)
	fmt.Println(r)
}

func TestNumIslands(t *testing.T) {
	gridArr := [...][]byte{{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'}}
	grid := make([][]byte, 0)
	for r := 0; r < len(gridArr); r++ {
		var cc []byte
		for c := 0; c < len(gridArr[0]); c++ {
			cc = append(cc, gridArr[r][c])
		}
		grid = append(grid, cc)
	}
	NumIslands(grid)
}

func TestLargestRectangleArea(t *testing.T) {
	heights := make([]int, 0)
	heights = append(heights, 2)
	heights = append(heights, 1)
	heights = append(heights, 5)
	heights = append(heights, 6)
	heights = append(heights, 2)
	heights = append(heights, 3)
	h := LargestRectangleArea(heights)
	fmt.Println(h)
}
