package stack_queue

import "testing"

func TestMyQueueConstructor(t *testing.T) {
	q := MyQueueConstructor()
	q.Push(3)
	q.Push(1)
	q.Push(2)
	q.Push(4)

	t.Log(q.Pop())
	t.Log(q.Empty())
	t.Log(q.Peek())
	t.Log(q.Peek())
	t.Log(q.Pop())
	t.Log(q.Pop())
	t.Log(q.Pop())
	t.Log(q.Empty())
	t.Log(q.Peek())
	t.Log(q.Pop())
}
