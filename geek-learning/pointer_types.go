package geek_learning

type Dog struct {
	name string
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

type Named interface {
	// Name 用于获取名字。
	Name() string
}
