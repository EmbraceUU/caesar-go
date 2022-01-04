package maps

import "testing"

func TestChangeMaps(t *testing.T) {
	data := make(map[string]interface{})
	key := "aaa"
	value := 0
	ChangeMaps(data, key, value)

	key2 := "bbb"
	value2 := "ccc"
	ChangeMaps(data, key2, value2)

	for k, v := range data {
		t.Log(k, v)
	}
}
