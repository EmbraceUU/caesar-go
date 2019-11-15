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

func TestSubStrArr(t *testing.T) {
	SubStrArr()
}

func TestNewCurElement(t *testing.T) {
	curListOld := []string{"btc", "eth", "eos", "win", "block", "bloc", "hot", "leo"}
	curListNew := []string{"btc", "eth", "eos", "win", "block", "bloc", "hot", "leo"}
	fmt.Println(NewCurElement(curListOld, curListNew))
	fmt.Println(len(NewCurElement(curListOld, curListNew)))
}
