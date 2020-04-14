package child

import "caesar-go/awesome/objects"

type One struct {
	objects.Base
}

func UpdateIN(aaa []int, aa int) {
	aaa[0] = aa
}
