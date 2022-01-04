package maps

// ChangeMaps 测试从函数里修改map, 是否会生效
func ChangeMaps(data map[string]interface{}, key string, value interface{}) {
	if data != nil {
		data[key] = value
	}
}
