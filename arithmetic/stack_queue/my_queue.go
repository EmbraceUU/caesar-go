package stack_queue

// MyQueue 请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）：
// 实现 MyQueue 类：
// void push(int x) 将元素 x 推到队列的末尾
// int pop() 从队列的开头移除并返回元素
// int peek() 返回队列开头的元素
// boolean empty() 如果队列为空，返回 true ；否则，返回 false
// https://leetcode-cn.com/problems/implement-queue-using-stacks/
type MyQueue struct {
	stack []int
	back  []int
}

//说明：
//
//你只能使用标准的栈操作 —— 也就是只有push to top,peek/pop from top,size, 和is empty操作是合法的。
//你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。

func MyQueueConstructor() MyQueue {
	return MyQueue{
		stack: make([]int, 0),
		back:  make([]int, 0),
	}
}

func (q *MyQueue) Push(x int) {
	// 如果back中有值, 先把back中的元素转入stack中
	for len(q.back) != 0 {
		val := q.back[len(q.back)-1]
		q.back = q.back[:len(q.back)-1]
		q.stack = append(q.stack, val)
	}
	q.stack = append(q.stack, x)
}

func (q *MyQueue) Pop() int {
	var x int
	if q.Empty() {
		return x
	}
	// 如果back不为空, 直接返回back的栈顶元素
	if len(q.back) > 0 {
		x = q.back[len(q.back)-1]
		q.back = q.back[:len(q.back)-1]
		return x
	}

	for len(q.stack) != 0 {
		x = q.stack[len(q.stack)-1]
		q.stack = q.stack[:len(q.stack)-1]
		q.back = append(q.back, x)
	}

	q.back = q.back[:len(q.back)-1]
	return x
}

func (q *MyQueue) Peek() int {
	var x int
	if q.Empty() {
		return x
	}
	// 如果back不为空, 直接返回back的栈顶元素
	if len(q.back) > 0 {
		x = q.back[len(q.back)-1]
		return x
	}

	for len(q.stack) != 0 {
		x = q.stack[len(q.stack)-1]
		q.stack = q.stack[:len(q.stack)-1]
		q.back = append(q.back, x)
	}

	return x
}

func (q *MyQueue) Empty() bool {
	return len(q.back) == 0 && len(q.stack) == 0
}
