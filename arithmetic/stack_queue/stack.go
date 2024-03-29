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

// CloneGraph 给你无向 连通 图中一个节点的引用，请你返回该图的 深拷贝（克隆）
//	输入：adjList = [[2,4],[1,3],[2,4],[1,3]]
//	输出：[[2,4],[1,3],[2,4],[1,3]]
// https://leetcode-cn.com/problems/clone-graph/
func CloneGraph(node *Node) *Node {
	cloneGraphVisited = make(map[*Node]*Node)
	return clone(node)
}

var cloneGraphVisited map[*Node]*Node

// 思路：使用递归的方法，并借用 map 来存放 clone 节点和 node 的映射关系
// 优化：将 map 提到外面作为全局变量, 减少了递归过程中传递参数带来的复制消耗
func clone(node *Node) *Node {
	if node == nil {
		return node
	}
	if c, ok := cloneGraphVisited[node]; ok {
		return c
	}

	cloneNode := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, 0),
	}
	cloneGraphVisited[node] = cloneNode
	for _, n := range node.Neighbors {
		cloneNode.Neighbors = append(cloneNode.Neighbors, clone(n))
	}
	return cloneNode
}

// NumIslands 求岛屿数量
// 思路: 使用DFS遍历岛屿，并将已经访问过的元素标记
// 当访问节点是岛屿时，深度遍历所有相邻节点，并标记，然后岛屿数量+1
// https://leetcode-cn.com/problems/number-of-islands/
func NumIslands(grid [][]byte) int {
	var dfs func(int, int)
	nr, nc := len(grid), len(grid[0])

	dfs = func(r, c int) {
		if r < 0 || r >= nr || c < 0 || c >= nc || grid[r][c] == '0' {
			return
		}

		grid[r][c] = '0'
		dfs(r+1, c)
		dfs(r, c+1)
		dfs(r-1, c)
		dfs(r, c-1)
	}

	var count int
	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if grid[r][c] == '1' {
				count++
				dfs(r, c)
			}
		}
	}
	return count
}

// LargestRectangleArea 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。
// 每个柱子彼此相邻，且宽度为 1 。 求在该柱状图中，能够勾勒出来的矩形的最大面积。
// https://leetcode-cn.com/problems/largest-rectangle-in-histogram/
func LargestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	stack := make([]int, 0)
	max := 0
	for i := 0; i <= len(heights); i++ {
		var cur int
		// 这里增加了一个元素, 即最右边加了一个0高度的元素
		if i == len(heights) {
			cur = 0
		} else {
			cur = heights[i]
		}

		// 当高度小于栈顶对应的元素高度时，出栈，从右向左计算
		for len(stack) != 0 && cur < heights[stack[len(stack)-1]] {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			h := heights[pop]

			w := i
			if len(stack) != 0 { // 栈里的值，肯定是比栈顶要小的，因为如果比栈顶大，就已经计算出矩形面积并且出栈了，所以直接拿当前栈顶，肯定是上一个栈顶的左侧严格小的下标
				peek := stack[len(stack)-1]
				// 计算出当前高度的最大宽度, 因为右边已经根据for循环的条件判断出比当前高度小了,
				// 而栈又是一个递增单调栈, 当前栈的栈顶肯定比当前高度低,
				// 所以i-peek-1是当前高度的从右到左的最大宽度
				w = i - peek - 1
			}
			if h*w > max {
				max = h * w
			}
		}

		stack = append(stack, i)
	}
	return max
}
