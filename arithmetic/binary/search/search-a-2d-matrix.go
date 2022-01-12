package search

// MatrixSearch 编写一个高效的算法来判断m * n矩阵中，是否存在一个目标值。该矩阵具有如下特性：
//	每行中的整数从左到右按升序排列。
//	每行的第一个整数大于前一行的最后一个整数。
// 输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
// 输出：true
// https://leetcode-cn.com/problems/search-a-2d-matrix
func MatrixSearch(matrix [][]int, target int) bool {
	// 先排除空数组的情况
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	// 定义声明row和col变量
	row, col := len(matrix), len(matrix[0])

	// 定义声明min和max，max按照一维数组处理
	min, max := 0, row*col-1

	for min+1 < max {
		mid := min + (max-min)/2

		// 二维数组的位置转换成一维数组
		// mid的位置应该是mid/col行，第mid%col列
		val := matrix[mid/col][mid%col]
		if val > target {
			max = mid
		} else if val < target {
			min = mid
		} else {
			return true
		}
	}

	// 当仅剩下min和max位置时，最后检查target值
	if matrix[min/col][min%col] == target || matrix[max/col][max%col] == target {
		return true
	}
	return false
}
