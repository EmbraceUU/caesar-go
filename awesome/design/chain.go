package design

import "fmt"

/*
链式调用模式
*/

// 链式调用模式, 类似与流式计算过程, 将每个处理过程相互解耦, 细分业务, 减少耦合度
// 并且易于扩展和调整业务逻辑
// 目前项目中有不少接口, 业务复杂, 数据结构庞大, 各个处理过程掺杂在一起, 代码可读性差, 既不容易扩展, 也不容易修改业务
// 往往需求变更时, 需要重新设计接口, 复用性也不高
// 如果切换成链式调用模式, 将Context作为公共数据, 或者返回结果, 在链上传递即可,
// 将相互独立的业务部分, 查分成组件, 在每个组件执行过程中对Context进行读写, 知道执行完毕
// 并且可以在Run方法中控制是否继续执行, 统一处理err, 增加日志
type ChainContext struct {
}

type ChainHandler interface {
	// 业务逻辑
	Do(c *ChainContext) error
	// 设置下一个组件
	SetNext(h ChainHandler) ChainHandler
	// 执行
	Run(c *ChainContext) error
}

// ChainNext 实现了 SetNext 方法, 为组件统一处理装配的过程
type ChainNext struct {
	nextHandler ChainHandler
}

// 将下一个组件装配到 ChainNext 中, 然后返回该组件, 继续向下装配
func (n *ChainNext) SetNext(h ChainHandler) ChainHandler {
	n.nextHandler = h
	return h
}

// 在 ChainNext 中实现了 Run 方法, 计算过程中统一执行此过程
// 将各个组件的共性部分抽离出来, 只保留差异化的业务处理过程
// 方便面向切面做处理
func (n *ChainNext) Run(c *ChainContext) error {
	if n.nextHandler != nil {
		if err := (n.nextHandler).Do(c); err != nil {
			return err
		}
		// 继续向下调用, 形成链式调用
		return (n.nextHandler).Run(c)
	}
	// 当没有下一个节点时, 说明已经全部执行完
	return nil
}

type NullHandler struct {
	ChainNext
}

func (h *NullHandler) Do(c *ChainContext) error {
	return nil
}

type BusinessHandler struct {
	ChainNext
}

func (h *BusinessHandler) Do(c *ChainContext) error {
	println("business 01 handler.")
	return nil
}

type Business02Handler struct {
	ChainNext
}

func (h *Business02Handler) Do(c *ChainContext) error {
	println("business 02 handler.")
	return nil
}

type Business03Handler struct {
	ChainNext
}

func (h *Business03Handler) Do(c *ChainContext) error {
	println("business 03 handler.")
	return nil
}

func ChainsAction() {
	nullHandler := &NullHandler{}

	// 将不同的业务组件装配到handler中
	nullHandler.
		SetNext(&BusinessHandler{}).
		SetNext(&Business02Handler{}).
		SetNext(&Business03Handler{})

	if err := nullHandler.Run(&ChainContext{}); err != nil {
		fmt.Println("Fail | Error:" + err.Error())
		return
	}
	fmt.Println("Success")
	return
}
