package http

import "testing"

func TestGetRequest(t *testing.T) {
	//GetRequest()
	maps := make(map[string]struct{})
	maps["aaa"] = struct{}{}
	delete(maps, "bbb")
}
