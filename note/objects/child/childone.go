package child

import "caesar-go/note/objects"

type ChildOne struct {
	objects.Base
}

func UpdateIN(aaa []int, aa int) {
	aaa[0] = aa
}
