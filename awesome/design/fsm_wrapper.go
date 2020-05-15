package design

import "fmt"

var (
	PowerOff   = FSMState("close")
	FirstGear  = FSMState("first G")
	SecondGear = FSMState("second G")
	ThirdGear  = FSMState("third G")

	PowerOffEvent   = FSMEvent("put close")
	FirstGearEvent  = FSMEvent("put first G")
	SecondGearEvent = FSMEvent("put second G")
	ThirdGearEvent  = FSMEvent("put third G")

	PowerOffHandler = FSMHandler(func() FSMState {
		fmt.Println("power off...")
		return PowerOff
	})
	FirstGearHandler = FSMHandler(func() FSMState {
		fmt.Println("turn on first gear...")
		return FirstGear
	})
	SecondGearHandler = FSMHandler(func() FSMState {
		fmt.Println("turn on second gear...")
		return SecondGear
	})
	ThirdGearEventHanlder = FSMHandler(func() FSMState {
		fmt.Println("turn on third gear...")
		return ThirdGear
	})
)

type FSMObject struct {
	*FSM
}

func NewFSMObject(initState FSMState) *FSMObject {
	return &FSMObject{
		FSM: NewFSM(initState),
	}
}

func FSMAction() {
	fsm := NewFSMObject(PowerOff)

	fsm.AddHandler(PowerOff, PowerOffEvent, PowerOffHandler)
	fsm.AddHandler(PowerOff, FirstGearEvent, FirstGearHandler)
	fsm.AddHandler(PowerOff, SecondGearEvent, SecondGearHandler)
	fsm.AddHandler(PowerOff, ThirdGearEvent, ThirdGearEventHanlder)

	fsm.AddHandler(FirstGear, PowerOffEvent, PowerOffHandler)
	fsm.AddHandler(FirstGear, FirstGearEvent, FirstGearHandler)
	fsm.AddHandler(FirstGear, SecondGearEvent, SecondGearHandler)
	fsm.AddHandler(FirstGear, ThirdGearEvent, ThirdGearEventHanlder)

	fsm.AddHandler(SecondGear, PowerOffEvent, PowerOffHandler)
	fsm.AddHandler(SecondGear, FirstGearEvent, FirstGearHandler)
	fsm.AddHandler(SecondGear, SecondGearEvent, SecondGearHandler)
	fsm.AddHandler(SecondGear, ThirdGearEvent, ThirdGearEventHanlder)

	fsm.AddHandler(ThirdGear, PowerOffEvent, PowerOffHandler)
	fsm.AddHandler(ThirdGear, SecondGearEvent, SecondGearHandler)
	fsm.AddHandler(ThirdGear, SecondGearEvent, SecondGearHandler)
	fsm.AddHandler(ThirdGear, ThirdGearEvent, ThirdGearEventHanlder)

	fsm.Call(ThirdGearEvent)
	fsm.Call(FirstGearEvent)
	fsm.Call(PowerOffEvent)
	fsm.Call(SecondGearEvent)
	fsm.Call(PowerOffEvent)

}
