package json

import "testing"

func TestMarshalIndent(t *testing.T) {
	type A struct {
		BB string `json:"bb"`
		CC string `json:"cc"`
		DD int    `json:"dd"`
	}

	var a A
	a.DD = 9
	a.CC = "cc"
	a.BB = "bb"

	result := MarshalIndent(a)
	t.Log(result)
}

func TestMarshal(t *testing.T) {
	type A struct {
		BB string `json:"bb"`
		CC string `json:"cc"`
		DD int    `json:"dd"`
	}

	var a A
	a.DD = 9
	a.CC = "cc"
	a.BB = "bb"

	result := Marshal(a)
	t.Log(result)
}
