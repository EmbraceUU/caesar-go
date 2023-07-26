package heap

import "testing"

func TestPriorityQueue_Pop(t *testing.T) {
	x := []int{7, 8, 2, 1, 5, 4, 3, 6}
	pq := Init(x)
	t.Log(pq.Pop())
	pq.Push(10)
	t.Log(pq.Pop())
	for {
		item, err := pq.Pop()
		if err != nil {
			break
		}
		t.Log(item)
	}
}
