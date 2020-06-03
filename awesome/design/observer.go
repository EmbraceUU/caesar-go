package design

import (
	"fmt"
	"reflect"
)

/*
学习观察者模式:
含义: 订阅者订阅主题, 主题通知订阅者
		不同的订阅者订阅不同的主题, 主题通知已经订阅的订阅者
优点: 高内聚, 将不同业务模块解耦, 易于扩展
模型:
1. 主题接口: Observable
	Attach() 添加订阅者
	Detach() 删除订阅者
	Notify() 通知订阅者
	observerList 观察者list
2. 观察者接口: ObserverInterface
	Do() 自身业务
 */

// todo 经典的观察者模式, 没有引用go本身经典的channel, 应该有改进的方式, 起码这个例子, not in go way

type Observable interface {
	// 增加观察者接口
	Attach(observerInterface ...ObserverInterface) Observable
	// 删除观察者接口
	Detach(observerInterface ObserverInterface) Observable
	// 通知接口
	Notify() error
}

// 观察者接口
type ObserverInterface interface {
	Do(observable Observable) error
}

// 具体的主题结构体
type ObservableConcrete struct {
	observerList []ObserverInterface
}

func (o *ObservableConcrete) Attach(observerInterface ...ObserverInterface) Observable {
	o.observerList = append(o.observerList, observerInterface...)
	return o
}

func (o *ObservableConcrete) Detach(observerInterface ObserverInterface) Observable {
	if o.observerList == nil || len(o.observerList) == 0 {
		return o
	}
	for k, item := range o.observerList {
		if observerInterface == item {
			fmt.Println(runFuncName(), " 注销:", reflect.TypeOf(observerInterface))
			o.observerList = append(o.observerList[:k], o.observerList[k+1:]...)
		}
	}
	return o
}

func (o *ObservableConcrete) Notify() error {
	if o.observerList == nil || len(o.observerList) == 0 {
		return nil
	}
	for _, item := range o.observerList {
		err := item.Do(o)
		if err != nil {
			fmt.Println(fmt.Errorf("%s 通知时报错: %v", runFuncName(), err))
			continue
		}
	}
	return nil
}

// 定义观察者
type Observer1 struct {}

func (o Observer1) Do(observable Observable) error {
	fmt.Println(runFuncName(), "observer1...")
	return nil
}

// 定义观察者2
type Observer2 struct {

}

func (o Observer2) Do(observable Observable) error {
	fmt.Println(runFuncName(), "observer2...")
	return nil
}

func ObserverAction() {
	// 定义主题
	// 定义观察者
	// 订阅主题
	// 通知观察者
	var observable ObservableConcrete
	var observer1 Observer1
	var observer2 Observer2

	observable.Attach(observer1, observer2)

	err := observable.Notify()
	if err != nil {
		fmt.Println(err)
	}
}