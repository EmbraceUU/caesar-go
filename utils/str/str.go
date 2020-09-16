package str

import (
	"regexp"
)

// 清除字符串中的数字
func Reg(key string) string {
	return regexp.MustCompile(`[0-9]+$`).ReplaceAllString(key, "")
}

// 清除字符串中, 从数字开始匹配的后面的字符
func RegSub(key string) string {
	return regexp.MustCompile(`[0-9]+.*$`).ReplaceAllString(key, "")
}
