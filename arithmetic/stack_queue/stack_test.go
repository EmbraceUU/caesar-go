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
	s := "xxx300[a]2[bc]"
	r := DecodeString(s)
	fmt.Println(r)
}
