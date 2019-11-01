package array

import (
	"fmt"
	"testing"
)

func TestArrDefineZero(t *testing.T) {
	arrStr := ArrDefineZero()
	fmt.Println(arrStr)
	fmt.Println(len(arrStr))
}

func TestArrDefineStructZero(t *testing.T) {
	arrStruc := ArrDefineStructZero()
	fmt.Println(arrStruc)
	fmt.Println(len(arrStruc))
}

func TestArrForIndex(t *testing.T) {
	ArrForIndex()
}

func TestMapArrStructPoint(t *testing.T) {
	MapArrStructPoint()
}
