package stack_queue

// UpdateMatrix 给定一个由 0 和 1 组成的矩阵 mat，
// 请输出一个大小相同的矩阵，其中每一个格子是 mat 中对应位置元素到最近的 0 的距离。
// https://leetcode-cn.com/problems/01-matrix/
func UpdateMatrix(mat [][]int) [][]int {
	// 多个起始点的BFS, 先把每个0入队列, 每个0就是起点
	q := make([][]int, 0)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 0 {
				q = append(q, []int{i, j})
			} else {
				// 设为-1, 用来标识未访问过的节点, 也是需要计算到0的距离的节点
				mat[i][j] = -1
			}
		}
	}

	// 定义一个方向数组, 用来指引当前节点的相邻节点
	directions := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

	for len(q) > 0 {
		// 出队
		point := q[0]
		q = q[1:]

		// 遍历四个方向
		for _, v := range directions {
			x := point[0] + v[0]
			y := point[1] + v[1]
			// 在mat范围内, 并且未曾访问过
			// 这里通过-1, 筛掉了作为起点的0值节点, 也筛掉了已经访问过的节点
			if x >= 0 && x < len(mat) && y >= 0 && y < len(mat[0]) && mat[x][y] == -1 {
				mat[x][y] = mat[point[0]][point[1]] + 1
				// 将当前的元素进队列，进行一次BFS
				q = append(q, []int{x, y})
			}
		}
	}

	// Note: 为什么这样做能标记出到0的最短距离 ？
	// 因为多个0节点, 是第一层访问的节点, 并且每次只访问相邻的节点, 并且是值为-1的节点, 也就是只访问下一层的节点
	// 那么对于下一层, 肯定是距离最近的, 因为距离为1
	// 然后再把这些访问到的节点入队, 等当前这层访问完再访问下一层
	// 下一层再去访问相邻的并且值为-1的节点, 那么距离肯定也是最小的, 然后赋值为当前层的距离+1
	return mat
}
