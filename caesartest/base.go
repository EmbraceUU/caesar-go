package caesar

import "fmt"

type Base struct {
	Name string
}

func (b *Base) InitBase() {
	fmt.Println("this is base")
}
