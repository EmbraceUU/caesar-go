package backtrack

func SolveNQueens(n int) [][]string {
	return solveNQueens(n)
}

// https://leetcode.cn/problems/n-queens/
// N皇后，需要暴力解法，使用回溯的思路，不断重试
func solveNQueens(n int) [][]string {
	var result [][]string
	record := make([]string, 0, n)
	// 记录每列选择情况，来判断竖列
	usedX := make([]bool, n)
	// 记录上一层的位置，来判断对角线
	lastIdx := -1
	bt2(&result, record, usedX, lastIdx, n)
	return result
}

func bt2(result *[][]string, record []string, usedX []bool, lastIdx int, n int) {
	// 判断是否结束
	// 一共n列，到达n列结束
	// 结束不一定正确
	if len(record) == len(usedX) {
		temp := make([]string, n)
		copy(temp, record)
		*result = append(*result, temp)
		return
	}

	// 遍历n遍
	for i := 0; i < n; i++ {
		// 	如果和自己是同一列过滤掉
		if usedX[i] {
			continue
		}
		//	如果和自己是斜对角过滤掉
		// 对角线不只是上一层，要将整条线上的都对比，所以需要一个 board 来记录每一行的记录才行
		if !isValid(record, i) {
			continue
		}

		//  保存记录
		usedX[i] = true
		record = append(record, getRecord(i, n))
		//	bt2
		bt2(result, record, usedX, i, n)
		//  删除记录
		usedX[i] = false
		record = record[:len(record)-1]
	}
}

func isValid(record []string, col int) bool {
	n := cap(record)
	row := len(record)

	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if record[i][j] == 'Q' {
			return false
		}
	}

	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if record[i][j] == 'Q' {
			return false
		}
	}
	return true
}

func getRecord(i int, n int) string {
	result := ""
	for j := 0; j < n; j++ {
		if i == j {
			result += "Q"
		} else {
			result += "."
		}
	}
	return result
}
