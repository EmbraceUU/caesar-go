package json

import "encoding/json"

// MarshalIndent 输出json格式的字符串
func MarshalIndent(source interface{}) string {
	brs, _ := json.MarshalIndent(source, "", " ")
	return string(brs)
}

// Marshal 输出json字符串
func Marshal(source interface{}) string {
	brs, _ := json.Marshal(source)
	return string(brs)
}
