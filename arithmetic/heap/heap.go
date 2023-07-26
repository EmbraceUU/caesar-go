package heap

import "fmt"

type PriorityQueue []int

func Init(x []int) PriorityQueue {
	pq := make(PriorityQueue, 0, len(x))
	for _, item := range x {
		pq.Push(item)
	}
	return pq
}

func (p *PriorityQueue) Push(x int) {
	*p = append(*p, x)
	p.heapInsert(x)
}

func (p *PriorityQueue) Pop() (int, error) {
	if len(*p) <= 0 {
		return 0, fmt.Errorf("empty queue. ")
	}
	old := *p
	// 交换根节点和末尾节点
	p.swap(0, len(old)-1)
	// 把换下来的最值返回
	x := old[len(old)-1]
	// 把换下来的最值的位置移出堆范围
	old = old[:len(old)-1]
	*p = old
	// 重新构建堆
	p.heapify()
	return x, nil
}

func (p *PriorityQueue) heapInsert(x int) {
	// 和自己的父节点比较，并尝试交换
	data := *p
	i := len(*p) - 1
	for {
		parent := (i - 1) / 2
		if parent == i || data[parent] < x {
			break
		}
		p.swap(parent, i)
		i = parent
	}
}

func (p *PriorityQueue) heapify() {
	// 和自己的左右子节点比较, 并尝试比较
	max := len(*p) - 1
	data := *p
	i := 0
	for {
		// 定位到左节点
		left := 2*i + 1
		// 判断是否溢出
		if left > max {
			break
		}

		// 判断最大值
		j := left
		// 判断右节点
		if left+1 <= max && data[left+1] < data[left] {
			j = left + 1
		}

		// 判断是否需要继续
		if data[i] < data[j] {
			break
		}

		// 交换，并向下继续
		p.swap(i, j)
		i = j
	}
}

func (p *PriorityQueue) swap(i, j int) {
	old := *p
	old[i], old[j] = old[j], old[i]
}
