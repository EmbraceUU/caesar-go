package array

// plusOne 加一
// 给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
// 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
// 你可以假设除了整数 0 之外，这个整数不会以零开头。
// 输入：digits = [4,3,2,1]
// 输出：[4,3,2,2]
// 【数组】【数学】
// https://leetcode-cn.com/problems/plus-one
func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return nil
	}

	if digits[0] == 9 {
		for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
			digits[i], digits[j] = digits[j], digits[i]
		}

		// [9,9,9,9]
		for i := 0; i < len(digits); i++ {
			digits[i]++
			if digits[i] >= 10 {
				digits[i] %= 10
				if i == len(digits)-1 {
					digits = append(digits, 1)
					break
				}
			} else {
				break
			}
		}

		for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
			digits[i], digits[j] = digits[j], digits[i]
		}
		return digits
	}

	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] >= 10 {
			digits[i] %= 10
		} else {
			break
		}
	}
	return digits
}
