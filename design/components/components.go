package components

import (
	"fmt"
	"reflect"
	"runtime"
)

type ComponentsContext struct{}

/*
将逻辑的执行组件话, 按照业务模块, 将相互独立的业务拆分成树形结构的组件

抽象出来的接口包含了Mount/Remove/Do方法
用一个BaseComponent结构体实现所有方法

子组件复用组合BaseComponent, 共用ChildComponent字段, 是共用字段, 不是共享这块内存

与责任链的区别是:
责任链是链式调用
组件的树形结构的调用方式
*/

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

type CheckoutPageComponent struct {
	BaseComponent
}

func (bc CheckoutPageComponent) Do(ctx *ComponentsContext) error {
	fmt.Println(runFuncName(), "订单结算页面组件...")

	// 执行子组件
	_ = bc.ChildDo(ctx)

	// 执行当前组件的逻辑

	return nil
}

type AddressComponent struct {
	BaseComponent
}

func (bc AddressComponent) Do(ctx *ComponentsContext) error {
	fmt.Println(runFuncName(), "地址组件...")

	// 执行子组件
	_ = bc.ChildDo(ctx)

	// 执行当前组件的逻辑

	return nil
}

// PayMethodComponent 支付方式组件
type PayMethodComponent struct {
	// 合成复用基础组件
	BaseComponent
}

// Do 执行组件&子组件
func (bc *PayMethodComponent) Do(ctx *ComponentsContext) (err error) {
	fmt.Println(runFuncName(), "支付方式组件...")

	// 执行子组件
	_ = bc.ChildDo(ctx)

	// 执行当前组件的逻辑
	return
}

// PromotionComponent 优惠信息组件
type PromotionComponent struct {
	// 合成复用基础组件
	BaseComponent
}

// Do 执行组件&子组件
func (bc *PromotionComponent) Do(ctx *ComponentsContext) (err error) {
	fmt.Println(runFuncName(), "优惠信息组件...")

	// 执行子组件
	_ = bc.ChildDo(ctx)

	// 执行当前组件的逻辑
	return
}

// AfterSaleComponent 售后组件
type AfterSaleComponent struct {
	// 合成复用基础组件
	BaseComponent
}

// Do 执行组件&子组件
func (bc *AfterSaleComponent) Do(ctx *ComponentsContext) (err error) {
	fmt.Println(runFuncName(), "售后组件...")

	// 执行子组件
	_ = bc.ChildDo(ctx)

	// 执行当前组件的逻辑
	return
}

func ComponentsAction() {
	// 初始化最大的组件
	checkoutPage := &CheckoutPageComponent{}

	// 初始化子组件
	address := &AddressComponent{}
	// 挂在子组件
	_ = address.Mount(&PromotionComponent{}, &AfterSaleComponent{})

	// 大组件下挂在子组件
	_ = checkoutPage.Mount(address, &PayMethodComponent{})

	// 执行
	_ = checkoutPage.Do(&ComponentsContext{})
}

// 获取正在运行的函数名
func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}
