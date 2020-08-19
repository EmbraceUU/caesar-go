package arithmetic

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	fmt.Println(Reverse(1563847412))
}

func TestSieve(t *testing.T) {
	Sieve()
}

func TestSieve01(t *testing.T) {
	Sieve01()
}
