package design

import (
	"fmt"
)

/*
模板模式, 其实在项目中用到的比较多, 将差异化的部分抽象出来, 然后将不变的逻辑和算法用一个实现类实现
比如交易所对接的接口, 每个交易所对接的接口是不同的, 可以把差异化的几口抽象出来, 每个交易所有自己不同的实现
将不变的逻辑和流程, 作为模板

应该说, 模板本身和抽象的接口以及具体的实现没有太大关系,
具体实现类型和模板的复用, 组成了完整的流程
*/

const (
	ConstActTypeTime   int32 = 1
	ConstActTypeTimes  int32 = 2
	ConstActTypeAmount int32 = 3
)

type TemplateContext struct{ *ActInfo }

type ActInfo struct {
	// 抽奖类型
	ActivityType int32
}

// 不同的实现差异, 抽象出来的接口
type BehaviorInterface interface {
	// 其他参数校验 不同模板实现方式不同
	checkParams(ctx *TemplateContext) error
	// 获取节点信息 不同模板实现方式不同
	getPrizesNode(ctx *TemplateContext) error
}

// 具体的实现
type TimeDraw struct {
	// 合成复用模板
	Lottery
}

func (d TimeDraw) checkParams(ctx *TemplateContext) error {
	fmt.Println(runFuncName(), "按时间抽奖类型:特殊参数校验...")
	return nil
}

func (d TimeDraw) getPrizesNode(ctx *TemplateContext) error {
	fmt.Println(runFuncName(), "do nothing(抽取该场次的奖品即可，无需其他逻辑)...")
	return nil
}

type TimesDraw struct {
	Lottery
}

func (d TimesDraw) checkParams(ctx *TemplateContext) error {
	fmt.Println(runFuncName(), "按抽奖次数抽奖类型:特殊参数校验...")
	return nil
}

func (d TimesDraw) getPrizesNode(ctx *TemplateContext) error {
	fmt.Println(runFuncName(), "1. 判断是该用户第几次抽奖...")
	fmt.Println(runFuncName(), "2. 获取对应node的奖品信息...")
	fmt.Println(runFuncName(), "3. 复写原所有奖品信息(抽取该node节点的奖品)...")
	return nil
}

type AmountDraw struct {
	Lottery
}

func (d AmountDraw) checkParams(ctx *TemplateContext) error {
	fmt.Println(runFuncName(), "按数额范围区间抽奖:特殊参数校验...")
	return nil
}

func (d AmountDraw) getPrizesNode(ctx *TemplateContext) error {
	fmt.Println(runFuncName(), "1. 判断属于哪个数额区间...")
	fmt.Println(runFuncName(), "2. 获取对应node的奖品信息...")
	fmt.Println(runFuncName(), "3. 复写原所有奖品信息(抽取该node节点的奖品)...")
	return nil
}

// 模板
type Lottery struct {
	b BehaviorInterface
}

// 固定不变的流程
func (lottery *Lottery) Run(ctx *TemplateContext) error {
	// 具体方法：校验活动编号(serial_no)是否存在、并获取活动信息
	if err := lottery.checkSerialNo(ctx); err != nil {
		return err
	}

	// ”抽象方法“：其他参数校验
	if err := lottery.checkParams(ctx); err != nil {
		return err
	}

	// ”抽象方法“：获取node奖品信息
	if err := lottery.getPrizesNode(ctx); err != nil {
		return err
	}

	// 具体方法：抽奖
	if err := lottery.drawPrizes(ctx); err != nil {
		return err
	}

	return nil
}

// 具体实现方法
func (lottery *Lottery) checkSerialNo(ctx *TemplateContext) error {
	fmt.Println(runFuncName(), "校验活动编号(serial_no)是否存在、并获取活动信息...")
	return nil
}

// 不同实现方法, 使用抽象接口调用
func (lottery *Lottery) checkParams(ctx *TemplateContext) error {
	return lottery.b.checkParams(ctx)
}

// 不同场景变化的算法 转化为依赖抽象
func (lottery *Lottery) getPrizesNode(ctx *TemplateContext) error {
	// 实际依赖的接口的抽象方法
	return lottery.b.getPrizesNode(ctx)
}

// drawPrizes 抽奖
func (lottery *Lottery) drawPrizes(ctx *TemplateContext) error {
	fmt.Println(runFuncName(), "抽奖...")
	return nil
}

// todo 为什么instance要将自身赋给包含的接口 ?
// todo 为什么实现类型中包含了Lottery, Lottery有包含了接口类型
func TemplateAction() {
	ctx := &TemplateContext{
		ActInfo: &ActInfo{
			ActivityType: ConstActTypeAmount,
		},
	}
	switch ctx.ActivityType {
	case ConstActTypeAmount:
		instance := &AmountDraw{}
		instance.b = instance
		_ = instance.Run(ctx)
	case ConstActTypeTimes:
		instance := &TimesDraw{}
		instance.b = instance
		_ = instance.Run(ctx)
	case ConstActTypeTime:
		instance := &TimeDraw{}
		instance.b = instance
		_ = instance.Run(ctx)
	}
}
