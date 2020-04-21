package design

import (
	"fmt"
	"reflect"
	"runtime"
)

type ComponentsContext struct{}

type Component interface {
	Mount(c Component, components ...Component) error
	Remove(c Component) error
	Do(cxt *ComponentsContext) error
}

type BaseComponent struct {
	ChildComponent []Component
}

func (bc *BaseComponent) Mount(c Component, components ...Component) error {
	bc.ChildComponent = append(bc.ChildComponent, c)
	if len(components) == 0 {
		return nil
	}
	bc.ChildComponent = append(bc.ChildComponent, components...)
	return nil
}

func (bc BaseComponent) Remove(c Component) error {
	if len(bc.ChildComponent) == 0 {
		return nil
	}
	for k, child := range bc.ChildComponent {
		if c == child {
			fmt.Println(runFuncName(), "移除:", reflect.TypeOf(child))
			bc.ChildComponent = append(bc.ChildComponent[:k], bc.ChildComponent[k+1:]...)
		}
	}
	return nil
}

func (bc BaseComponent) Do(ctx *ComponentsContext) error {
	return nil
}

func (bc BaseComponent) ChildDo(ctx *ComponentsContext) error {
	// 这里是一个链式的逻辑, 如果出现错误, 直接退出了
	for _, child := range bc.ChildComponent {
		if err := child.Do(ctx); err != nil {
			return err
		}
	}
	return nil
}

// 获取正在运行的函数名
func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}
