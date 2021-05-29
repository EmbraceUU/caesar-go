package stack_queue

import "testing"

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
