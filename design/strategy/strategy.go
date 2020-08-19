package strategy

/*
学习策略模式:
含义: 不同的算法按照统一的标准封装(肯定用到了interface), 根据不同场景(参数), 决策(根据参数做选择)使用不同的算法
优点: 高内聚, 松耦合
模型:
标准: 定义一个接口
策略: 定义不同的实现
决策: 根据type不同, 采用不同的实现
*/

// todo gorm在初始化客户端的时候, 是否也是一种策略模式, 接口本身只是告诉gorm采用的dirver类型以及通用的url, 并统一返回了已经初始化的DB

const (
	Strategy1 = "s1"
	Strategy2 = "s2"
	Strategy3 = "s3"
)

type StrategyContext struct {
	Type string
}

// 标准即是接口
type IStrategy interface {
	Action1(c *StrategyContext) error
	Action2(c *StrategyContext) error
}

type ImpStrategy struct {
}

func (i ImpStrategy) Action1(c *StrategyContext) error {
	return nil
}

func (i ImpStrategy) Action2(c *StrategyContext) error {
	return nil
}

type ImpStrategy2 struct {
}

func (i ImpStrategy2) Action1(c *StrategyContext) error {
	return nil
}

func (i ImpStrategy2) Action2(c *StrategyContext) error {
	return nil
}

func StrategyAction() {
	ctx := &StrategyContext{
		Type: Strategy1,
	}
	var instance IStrategy
	switch ctx.Type {
	case Strategy1:
		instance = &ImpStrategy{}
	case Strategy2:
		instance = &ImpStrategy2{}
	default:
		panic("type is nil")
	}
	_ = instance.Action1(ctx)
	_ = instance.Action2(ctx)
}
