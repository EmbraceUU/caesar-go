package channel

var (
	limitChan chan struct{}
)

func Channel() {
	limitChan = make(chan struct{}, 2)

}
