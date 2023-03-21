package order

import (
	fsm "github.com/WTongStudio/fsm-order-demo"
	"github.com/WTongStudio/fsm-order-demo/order/action"
	"github.com/WTongStudio/fsm-order-demo/order/processor"
)

// | 状态            | 编码 | 允许操作                                        |
// | -------------- | ---- | --------------------------------------------- |
// | 待支付          | 0    | 支付: 待确认; 取消: 已取消; 支付确认: 已支付        |
// | 已取消          | 1    | 无                                             |
// | 待确认          | 2    | 支付确认: 待发货                                 |
// | 待发货          | 3    | 发货: 待收货; 申请退款: 售后中-退款                |
// | 售后中-退款      | 4    | 取消售后: 待发货; 退款完成: 已完成                 |
// | 待收货          | 5    | 签收: 已签收                                    |
// | 已签收          | 6    | 申请退货退款: 售后中-退货退款; 订单完成: 已完成      |
// | 售后中-退货退款  | 7    | 取消售后: 已签收; 退款完成: 已完成                 |
// | 已完成          | 8    | 无                                             |

// 状态
const (
	StateWaitPay     = iota // 待支付
	StateCancel             // 已取消
	StateWaitConfirm        // 待确认
	StateWaitDeliver        // 待发货
	StateRefund             // 售后中-退款
	StateWaitReceive        // 待收货
	StateSigned             // 已签收
	StateGoodsRefund        // 售后中-退货退款
	StateCompleted          // 已完成
)

// StateDesc 状态描述
var StateDesc = map[fsm.State]string{
	StateWaitPay:     "待支付",
	StateCancel:      "已取消",
	StateWaitConfirm: "待确认",
	StateWaitDeliver: "待发货",
	StateRefund:      "售后中-退款",
	StateWaitReceive: "待收货",
	StateSigned:      "已签收",
	StateGoodsRefund: "售后中-退货退款",
	StateCompleted:   "已完成",
}

// 事件
const (
	EventPay              = "pay"                // 支付事件
	EventCancel           = "cancel"             // 取消事件
	EventPayConfirm       = "pay_confirm"        // 支付确认事件
	EventDeliver          = "deliver"            // 发货事件
	EventApplyRefund      = "apply_refund"       // 申请退款事件
	EventCancelRefund     = "cancel_refund"      // 取消售后事件
	EventRefundCompleted  = "refund_completed"   // 退款完成事件
	EventSigned           = "signed"             // 签收事件
	EventApplyGoodsRefund = "apply_goods_refund" // 申请退货退款事件
	EventCompleted        = "completed"          // 订单完成事件
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
		// 支付确认事件：待支付 ---> 待发货
		EventPayConfirm: fsm.Transition{
			From: StateWaitPay, Event: EventPayConfirm, To: StateWaitDeliver, Action: action.PayConfirm, Processor: nil,
		},
	},
	StateWaitConfirm: {
		// 支付确认事件：待确认 ---> 待发货
		EventPayConfirm: fsm.Transition{
			From: StateWaitConfirm, Event: EventPayConfirm, To: StateWaitDeliver, Action: action.PayConfirm,
			Processor: nil,
		},
	},
	StateWaitDeliver: {
		// 发货事件：待发货 ---> 待收货
		EventDeliver: fsm.Transition{
			From: StateWaitDeliver, Event: EventDeliver, To: StateWaitReceive, Action: action.Delive, Processor: nil,
		},
		// 申请退款事件：待发货 ---> 售后中-退款
		EventApplyRefund: fsm.Transition{
			From: StateWaitDeliver, Event: EventApplyRefund, To: StateRefund, Action: action.ApplyRefund, Processor: nil,
		},
	},
	StateRefund: {
		// 取消售后事件：售后中-退款 ---> 待发货
		EventCancelRefund: fsm.Transition{
			From: StateRefund, Event: EventCancelRefund, To: StateWaitDeliver, Action: action.CancelRefund,
			Processor: nil,
		},
		// 退款完成事件：售后中-退款 ---> 已完成
		EventRefundCompleted: fsm.Transition{
			From: StateRefund, Event: EventRefundCompleted, To: StateCompleted, Action: action.RefundCompleted,
			Processor: nil,
		},
	},
	StateWaitReceive: {
		// 签收事件：待收货 ---> 已签收
		EventSigned: fsm.Transition{
			From: StateWaitReceive, Event: EventSigned, To: StateSigned, Action: action.Signed, Processor: nil,
		},
	},
	StateSigned: {
		// 申请退货退款事件：已签收 ---> 售后中-退货退款
		EventApplyGoodsRefund: fsm.Transition{
			From: StateSigned, Event: EventApplyGoodsRefund, To: StateGoodsRefund, Action: action.ApplyGoodsRefund,
			Processor: nil,
		},
		// 订单完成事件：已签收 ---> 已完成
		EventCompleted: fsm.Transition{
			From: StateSigned, Event: EventCompleted, To: StateCompleted, Action: action.Completed, Processor: nil,
		},
	},
	StateGoodsRefund: {
		// 取消售后事件：售后中-退货退款 ---> 已签收
		EventCancelRefund: fsm.Transition{
			From: StateGoodsRefund, Event: EventCancelRefund, To: StateSigned, Action: action.CancelRefund,
			Processor: nil,
		},
		// 退款完成事件：售后中-退货退款 ---> 已完成
		EventRefundCompleted: fsm.Transition{
			From: StateGoodsRefund, Event: EventRefundCompleted, To: StateCompleted, Action: action.RefundCompleted,
			Processor: nil,
		},
	},
}

// NewStateMachine 生成状态机
func NewStateMachine() *fsm.StateMachine {
	sm := &fsm.StateMachine{
		Processor: &processor.EventProcessor{},
		Graph:     &fsm.StateGraph{},
	}
	sm.SetName("子订单状态图表")
	sm.SetStart(StateWaitPay)
	sm.SetEnd(StateCompleted)
	sm.SetStates(StateDesc)
	sm.SetTransitions(transitions)
	return sm
}
