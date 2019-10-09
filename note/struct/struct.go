package _struct

type BoolStruct struct {
	Bool1 bool `json:"bool_1"`
	Bool2 bool `json:"bool_2"`
}

func GetBoolStruct() BoolStruct {
	var Bs BoolStruct
	return Bs
}
