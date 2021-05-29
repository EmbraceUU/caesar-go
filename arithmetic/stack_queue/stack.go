package stack_queue

// 栈的特点是 FILO 后进先出
// 常用于临时保存一些数据，DFS 遍历

// MinStack 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
// push(x) —— 将元素 x 推入栈中。
// pop() —— 删除栈顶的元素。
// top() —— 获取栈顶元素。
// getMin() —— 检索栈中的最小元素。
type MinStack struct {
	data []int
	min  []int
}

// Constructor initialize your data structure here.
func Constructor() MinStack {
	return MinStack{
		data: make([]int, 0),
		min:  make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)

	// 获取到当前最小值，将val和m的最小值放到min的顶部
	m := this.GetMin()
	if m < val {
		this.min = append(this.min, m)
	} else {
		this.min = append(this.min, val)
	}
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-1]
	// TODO 这里是重点, 意味着 min 记录的是每次 Push 以后，data 里面当时的最小值
	// 所以当发生 Pop 时，min 也删除这个值在 Push 时 data 内的最小值
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.min) == 0 {
		return 1 << 31
	}
	return this.min[len(this.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
