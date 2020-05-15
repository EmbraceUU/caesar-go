package design

import (
	"fmt"
	"sync"
)

/*
状态机模式
*/

type FSMState string            // 状态
type FSMEvent string            // 事件
type FSMHandler func() FSMState // 处理方法 返回新的状态

type FSM struct {
	mu       sync.Mutex
	state    FSMState                             // 当前状态
	handlers map[FSMState]map[FSMEvent]FSMHandler // 处理集 每个状态可以触发有限个事件 执行有限个处理
}

func (f *FSM) getState() FSMState {
	return f.state
}

func (f *FSM) setState(state FSMState) {
	f.state = state
}

// 给某个状态添加处理方法
func (f *FSM) AddHandler(state FSMState, event FSMEvent, handler FSMHandler) *FSM {
	if _, ok := f.handlers[state]; !ok {
		f.handlers[state] = make(map[FSMEvent]FSMHandler)
	}
	if _, ok := f.handlers[state][event]; ok {
		fmt.Printf("WARN state %s event %s is defined", state, event)
	}
	f.handlers[state][event] = handler
	return f
}

// 事件处理
func (f *FSM) Call(event FSMEvent) FSMState {
	f.mu.Lock()
	defer f.mu.Unlock()
	events := f.handlers[f.getState()]
	if events == nil {
		return f.getState()
	}
	if fn, ok := events[event]; ok {
		oldState := f.getState()
		f.setState(fn())
		newState := f.getState()
		fmt.Println("state change from ", oldState, " to ", newState)
	}
	return f.getState()
}

func NewFSM(initState FSMState) *FSM {
	return &FSM{
		state:    initState,
		handlers: make(map[FSMState]map[FSMEvent]FSMHandler),
	}
}
