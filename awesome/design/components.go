package design

import "fmt"

/*
组合模式
*/
type Context struct {
}

type Handler interface {
	// 业务逻辑
	Do(c *Context) error
	// 设置下一个组件
	SetNext(h Handler) Handler
	// 执行
	Run(c *Context) error
}

type Next struct {
	nextHandler Handler
}

func (n *Next) SetNext(h Handler) Handler {
	n.nextHandler = h
	return h
}

func (n *Next) Run(c *Context) error {
	if n.nextHandler != nil {
		if err := (n.nextHandler).Do(c); err != nil {
			return err
		}
		return (n.nextHandler).Run(c)
	}
	return nil
}

type NullHandler struct {
	Next
}

func (h *NullHandler) Do(c *Context) error {
	return nil
}

type BusiesHandler struct {
	Next
}

func (h *BusiesHandler) Do(c *Context) error {
	println("busies handler.")
	return nil
}

func Components() {
	nullHandler := &NullHandler{}

	nullHandler.SetNext(&BusiesHandler{})

	if err := nullHandler.Run(&Context{}); err != nil {
		fmt.Println("Fail | Error:" + err.Error())
		return
	}
	fmt.Println("Success")
	return
}
