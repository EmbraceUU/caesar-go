package graph

// Node 图中的节点
type Node struct {
	Val       int
	Neighbors []*Node
}

func NewNode(data [][]int) *Node {
	var nodes []*Node
	for i := range data {
		node := new(Node)
		node.Val = i + 1
		nodes = append(nodes, node)
	}

	for i, v := range data {
		for _, vv := range v {
			nodes[i].Neighbors = append(nodes[i].Neighbors, nodes[vv-1])
		}
	}
	return nodes[0]
}
