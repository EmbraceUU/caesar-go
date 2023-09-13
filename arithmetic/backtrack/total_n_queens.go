package backtrack

func totalNQueens(n int) int {
	var result int
	record := make([]string, 0, n)
	// 记录每列选择情况，来判断竖列
	usedX := make([]bool, n)
	// 记录上一层的位置，来判断对角线
	lastIdx := -1
	bt3(&result, record, usedX, lastIdx, n)
	return result
}

func bt3(result *int, record []string, usedX []bool, lastIdx int, n int) {
	// 判断是否结束
	// 一共n列，到达n列结束
	// 结束不一定正确
	if len(record) == len(usedX) {
		*result++
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
		bt3(result, record, usedX, i, n)
		//  删除记录
		usedX[i] = false
		record = record[:len(record)-1]
	}
}
