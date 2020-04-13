package _map

import (
	"sync"
)

var (
	TMap map[string]*AAA
	TMm  = new(sync.RWMutex)
)

// 使用指针的话, 只获得指针, 然后进行操作, 不用重新再赋值 ?
type AAA struct {
	D map[string]string
}

func WriteMap(key string, data *AAA) {
	TMap[key] = data
}

// awesome
// 如果用map作为cache, 那么返回数据时不要直接将map返回
// 应该生成一个新的map返回出去,
// 因为如果外部的调用方进行delete操作, 是直接操作的cache map

func InitMap() {
	TMap = make(map[string]*AAA)
}

func ReadMap(key string) *AAA {
	TMm.RLock()
	defer TMm.RUnlock()

	item, _ := TMap[key]
	//res := make(map[string]string)
	//if ok {
	//    // 这样赋值, 返回的map在发生改变后, 依然会影响TMap, 这说明赋值的是指针, 而非值
	//    res = item
	//    return res
	//}
	return item

	// 做值赋值, 不会受TMap的影响
	//res := make(map[string]string)
	//if ok {
	//	for k, v := range item {
	//		res[k] = v
	//	}
	//	return res
	//}
	//return nil
}
