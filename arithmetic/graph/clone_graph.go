package graph

// CloneGraph 给你无向 连通 图中一个节点的引用，请你返回该图的 深拷贝（克隆）。
// 图中的每个节点都包含它的值 val（int） 和其邻居的列表（list[Node]）。
//
// 测试用例格式：
//
// 简单起见，每个节点的值都和它的索引相同。
// 例如，第一个节点值为 1（val = 1），第二个节点值为 2（val = 2），以此类推。该图在测试用例中使用邻接列表表示。
//
// 邻接列表 是用于表示有限图的无序列表的集合。每个列表都描述了图中节点的邻居集。
//
// 给定节点将始终是图中的第一个节点（值为 1）。你必须将给定节点的拷贝作为对克隆图的引用返回。
// https://leetcode-cn.com/problems/clone-graph
func CloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	// 使用BFS或者DFS

	queue := make([]*Node, 0)
	set := make(map[*Node]*Node)

	queue = append(queue, node)
	// 需要在进队列之前copy节点，否则无法确定第一个节点的位置
	cp := new(Node)
	cp.Val = node.Val
	set[node] = cp

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, v := range cur.Neighbors {
			// 如果已经存在了，不需要进队列，但要把对应的copy节点同步到当前节点的copy节点的邻接点中
			if cpv, ok := set[v]; ok {
				set[cur].Neighbors = append(set[cur].Neighbors, cpv)
				continue
			}

			// 未遍历到的节点进队列，同时copy一份，并将其同步到copy节点的邻接点中
			queue = append(queue, v)
			cpv := new(Node)
			cpv.Val = v.Val
			set[v] = cpv
			set[cur].Neighbors = append(set[cur].Neighbors, cpv)
		}
	}

	return cp
}

// CloneGraphII DFS stack 做法
func CloneGraphII(node *Node) *Node {
	if node == nil {
		return nil
	}

	// 尝试 DFS
	// DFS 使用到了栈
	// 出栈后，找到一个邻接点，将当前节点和邻接点再次入栈
	stack := make([]*Node, 0)
	set := make(map[*Node]*Node)
	//[[2,4],[1,3],[2,4],[1,3]]
	stack = append(stack, node)
	cp := new(Node)
	cp.Val = node.Val
	set[node] = cp

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for len(cur.Neighbors) > 0 {
			// FIXME: 这里会修改原node的数据，导致无法校验结果
			// 使用递归可以避免这种问题，因为递归在逻辑上是一次性同步完节点的所有邻接点
			// 而非递归需要不断的压栈和入栈，导致重复遍历到某些邻接点
			// 除非牺牲空间或者时间去做去重，否则没有递归好。
			v := cur.Neighbors[0]
			cur.Neighbors = cur.Neighbors[1:]

			// TAG: 01
			if cpv, ok := set[v]; ok {
				set[cur].Neighbors = append(set[cur].Neighbors, cpv)
				continue
			}

			stack = append(stack, cur)
			stack = append(stack, v)
			cpv := new(Node)
			cpv.Val = v.Val
			set[v] = cpv
			set[cur].Neighbors = append(set[cur].Neighbors, cpv)
			break
		}
	}

	return cp
}

func CloneGraphIII(node *Node) *Node {
	if node == nil {
		return node
	}

	set := make(map[*Node]*Node)

	var clone func(n *Node) *Node
	clone = func(n *Node) *Node {

		if cp, ok := set[n]; ok {
			return cp
		}

		cp := new(Node)
		cp.Val = n.Val
		set[n] = cp

		for _, v := range n.Neighbors {
			cp.Neighbors = append(cp.Neighbors, clone(v))
		}

		return cp
	}

	return clone(node)
}
