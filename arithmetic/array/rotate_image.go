package array

// rotateImage 旋转图像
// 给定一个 n×n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
// 你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
// 输入：matrix = [[1(0,0),2(0,1),3(0,2)],
//                [4,5,6],
//                [7,8,9]]
// 输出：[[7,4,1(0,2)],
//       [8,5,2(1,2)],
//       [9,6,3(2,2)]]
// 【数组】【矩阵】
// https://leetcode-cn.com/problems/rotate-image
func rotateImage(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] =
				matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
		}
	}
}
