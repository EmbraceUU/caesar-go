package links

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	// 借助于小根堆
	// pop出来一个node，然后加入到list中
	// 并且将pop出的node的next push进堆中
	// 直到堆中没有数据

	// 将每个list的第一个note加入到小根堆中
	pq := make(PriorityQueue, 0, len(lists))
	for _, l := range lists {
		if l == nil {
			continue
		}
		pq.Push(l)
	}

	dummy := new(ListNode)
	p := dummy
	for {
		node, err := pq.Pop()
		if err != nil {
			break
		}
		p.Next = node
		p = p.Next

		newNode := node.Next
		if newNode != nil {
			pq.Push(newNode)
		}
	}
	return dummy.Next
}

type PriorityQueue []*ListNode

func (p *PriorityQueue) Push(x *ListNode) {
	*p = append(*p, x)
	p.heapInsert(x)
}

func (p *PriorityQueue) Pop() (*ListNode, error) {
	if len(*p) <= 0 {
		return nil, fmt.Errorf("empty queue. ")
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

func (p *PriorityQueue) heapInsert(x *ListNode) {
	// 和自己的父节点比较，并尝试交换
	data := *p
	i := len(*p) - 1
	for {
		parent := (i - 1) / 2
		if parent == i || data[parent].Val < x.Val {
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
		if left+1 <= max && data[left+1].Val < data[left].Val {
			j = left + 1
		}

		// 判断是否需要继续
		if data[i].Val < data[j].Val {
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
