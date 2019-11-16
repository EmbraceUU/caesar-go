package page

import (
	"fmt"
	"testing"
)

func TestPagingList(t *testing.T) {
	arr := []interface{}{"1", "2", "3", "4", "5", "6"}
	fmt.Print(PagingList(arr, "desc", 3, 4))
}
