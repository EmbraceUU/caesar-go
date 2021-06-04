package stack_queue

import "strconv"

// 栈的特点是 FILO 后进先出
// 常用于临时保存一些数据，DFS 遍历

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

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

// ********************************** //

// EvalRPN 逆波兰表达式求值
// 思路: ["4","13","5","/","+"] -> (4 + (13 / 5))
// 1. 利用 stack, 将数组压栈, 遇到运算符时, 出栈计算, 然后再压栈
// https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/
func EvalRPN(tokens []string) int {
	if tokens == nil || len(tokens) == 0 {
		return 0
	}

	stack := make([]int, 0)
	for i := 0; i < len(tokens); i++ {
		v := tokens[i]
		// 如果是数值, 转换后压栈
		if v != "+" && v != "-" && v != "*" && v != "/" {
			num, _ := strconv.Atoi(v)
			stack = append(stack, num)
			continue
		}

		num1 := stack[len(stack)-1]
		num2 := stack[len(stack)-2]
		stack = stack[:len(stack)-2]

		switch v {
		case "+":
			stack = append(stack, num2+num1)
		case "-":
			stack = append(stack, num2-num1)
		case "*":
			stack = append(stack, num2*num1)
		case "/":
			stack = append(stack, num2/num1)
		}
	}

	return stack[0]
}

// DecodeString 给定一个经过编码的字符串，返回它解码后的字符串。
// s = "3[a2[c]]"  -> "accaccacc"
// 思路: 利用栈的特性, 遍历字符, 以此压栈, 当遇到 ], 以此出栈遇到第一个 [, 然后中间的字符为需要循环的内容, 再出栈的是循环的次数
// TODO 需要优化, 内存使用太高
// https://leetcode-cn.com/problems/decode-string/
func DecodeString(s string) string {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == ']' {
			// 遍历得到需要循环的内容
			temp := make([]byte, 0)
			j := 0
			for len(stack) > 0 {
				j++
				c := stack[len(stack)-j]
				if c != '[' {
					temp = append(temp, c)
				} else {
					break
				}
			}

			// 完成字符的出栈操作
			// TODO j在上面多加了一次1, 所以在这里不需要再单独将 [ 出栈
			stack = stack[:len(stack)-j]

			idx := 1
			for len(stack) >= idx && stack[len(stack)-idx] >= '0' && stack[len(stack)-idx] <= '9' {
				idx++
			}

			num, _ := strconv.Atoi(string(stack[len(stack)-idx+1:]))
			stack = stack[:len(stack)-idx+1]

			// TODO []byte可以直接转换为string
			// TODO byte可以直接和字符比较
			// TODO 循环时需将要循环的内容提前赋给临时变量, 要不然会直接翻倍
			// 再把 cycle 压栈
			for k := 0; k < num; k++ {
				for m := len(temp) - 1; m >= 0; m-- {
					stack = append(stack, temp[m])
				}
			}
			continue
		}
		stack = append(stack, s[i])
	}

	return string(stack)
}

// InorderTraversal 二叉树的中序遍历, 利用 stack
// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
func InorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		// root 不为空的时候 压栈
		for root != nil {
			stack = append(stack, root)
			// 跳转到左子树
			root = root.Left
		}

		// 通过 node 访问值, 而通过 root 往 stack 里面压 node
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)

		// 跳转到右子树
		root = node.Right
	}
	return result
}
