package main_order

import (
	fsm "github.com/WTongStudio/fsm-order-demo"
	"github.com/WTongStudio/fsm-order-demo/main_order/action"
	"github.com/WTongStudio/fsm-order-demo/main_order/processor"
)

// | 状态   | 编码 | 允许操作及目标状态                          |
// | ----- | ---- | ---------------------------------------- |
// | 待支付 | 0    | 支付: 待确认; 取消: 已取消; 支付确认: 已支付   |
// | 已取消 | 1    | 无                                        |
// | 待确认 | 2    | 支付确认: 已支付                            |
// | 已支付 | 3    | 无                                        |

// 状态
const (
	StateWaitPay     = iota // 待支付
	StateCancel             // 已取消
	StateWaitConfirm        // 待确认
	StatePaid               // 已支付
)

// StateDesc 状态描述
var StateDesc = map[fsm.State]string{
	StateWaitPay:     "待支付",
	StateCancel:      "已取消",
	StateWaitConfirm: "待确认",
	StatePaid:        "已支付",
}

// 事件
const (
	EventPay        = "pay"         // 支付事件
	EventPayConfirm = "pay_confirm" // 支付确认事件
	EventCancel     = "cancel"      // 取消事件
)

// transitions 转变器
var transitions = map[fsm.State]map[fsm.Event]fsm.Transition{
	StateWaitPay: {
		// 取消事件：待支付 ---> 已取消
		EventCancel: fsm.Transition{
			From: StateWaitPay, Event: EventCancel, To: StateCancel, Action: action.Cancel, Processor: nil,
		},
		// 支付事件：待支付 ---> 待确认
		EventPay: fsm.Transition{
			From: StateWaitPay, Event: EventPay, To: StateWaitConfirm, Action: action.Pay,
			Processor: &processor.PayEventProcessor{},
		},
		// 支付确认事件：待支付 ---> 已支付
		EventPayConfirm: fsm.Transition{
			From: StateWaitPay, Event: EventPayConfirm, To: StatePaid, Action: action.PayConfirm, Processor: nil,
		},
	},
	StateWaitConfirm: {
		// 支付确认事件：待确认 ---> 已支付
		EventPayConfirm: fsm.Transition{
			From: StateWaitConfirm, Event: EventPayConfirm, To: StatePaid, Action: action.PayConfirm, Processor: nil,
		},
	},
}

// NewStateMachine 生成状态机
func NewStateMachine() *fsm.StateMachine {
	sm := &fsm.StateMachine{
		Processor: &processor.EventProcessor{},
		Graph:     &fsm.StateGraph{},
	}
	sm.SetName("主订单状态图表")
	sm.SetStart(StateWaitPay)
	sm.SetEnd(StatePaid)
	sm.SetStates(StateDesc)
	sm.SetTransitions(transitions)
	return sm
}
