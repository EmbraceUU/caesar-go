package _map

func TestMap(curMap map[string]string) {
	for key, value := range curMap {
		value = "123123"
		curMap[key] = value
	}
}

func WriteMap(filed, key string, value float64) map[string]map[string]float64 {
	diMap := make(map[string]map[string]float64)
	diMap[filed][key] = value
	return diMap
}

// note
// 如果用map作为cache, 那么返回数据时不要直接将map返回
// 应该生成一个新的map返回出去,
// 因为如果外部的调用方进行delete操作, 是直接操作的cache map
