package refund

import (
	fsm "github.com/WTongStudio/fsm-order-demo"
	"github.com/WTongStudio/fsm-order-demo/refund/action"
	"github.com/WTongStudio/fsm-order-demo/refund/processor"
)

// | 状态   | 编码 | 允许操作                                                               |
// | ----- | ---- | --------------------------------------------------------------------- |
// | 待审批 | 0    | 通过: 已通过; 驳回: 已驳回; 取消：已取消                                   |
// | 已取消 | 1    | 无                                                                    |
// | 已驳回 | 2    | 无                                                                    |
// | 已通过 | 3    | 提交退款申请(未发货订单): 退款中; 等待用户寄回(已发货订单): 退货中; 取消：已取消 |
// | 退货中 | 4    | 发货: 待收货; 取消：已取消                                               |
// | 待收货 | 5    | 签收: 退款中                                                           |
// | 退款中 | 6    | 退款完成: 已完成                                                        |
// | 已完成 | 7    | 无                                                                    |

const (
	StateWaitApprove = iota // 待审批
	StateCancel             // 已取消
	StateRefused            // 已驳回
	StateAgreed             // 已通过
	StateWaitDeliver        // 退货中
	StateWaitReceive        // 待收货
	StateWaitRefund         // 退款中
	StateCompleted          // 已完成
)

// StateDesc 订单描述
var StateDesc = map[fsm.State]string{
	StateWaitApprove: "待审批",
	StateCancel:      "已取消",
	StateRefused:     "已驳回",
	StateAgreed:      "已通过",
	StateWaitDeliver: "退货中",
	StateWaitReceive: "待收货",
	StateWaitRefund:  "退款中",
	StateCompleted:   "已完成",
}

const (
	EventCancel          = "cancel"           // 取消事件
	EventRefuse          = "refuse"           // 驳回事件
	EventAgree           = "agree"            // 通过事件
	EventRefund          = "refund"           // 提交退款申请事件(未发货订单)
	EventGoodsRefund     = "goods_refund"     // 等待用户寄回事件(已发货订单)
	EventDeliver         = "deliver"          // 发货事件
	EventSigned          = "signed"           // 签收事件
	EventRefundCompleted = "refund_completed" // 退款完成事件
)

// transitions 转变器
var transitions = map[fsm.State]map[fsm.Event]fsm.Transition{
	StateWaitApprove: {
		// 取消事件：待审批 ---> 已取消
		EventCancel: fsm.Transition{
			From: StateWaitApprove, Event: EventCancel, To: StateCancel, Action: action.Cancel, Processor: nil,
		},
		// 驳回事件：待审批 ---> 已驳回
		EventRefuse: fsm.Transition{
			From: StateWaitApprove, Event: EventRefuse, To: StateRefused, Action: action.Refuse, Processor: nil,
		},
		// 通过事件：待审批 ---> 已通过
		EventAgree: fsm.Transition{
			From: StateWaitApprove, Event: EventAgree, To: StateAgreed, Action: action.Agree, Processor: nil,
		},
	},
	StateAgreed: {
		// 提交退款申请事件(未发货订单)：已通过 ---> 退款中
		EventRefund: fsm.Transition{
			From: StateAgreed, Event: EventRefund, To: StateWaitRefund, Action: action.Refund, Processor: nil,
		},
		// 等待用户寄回事件(已发货订单)：已通过 ---> 退货中
		EventGoodsRefund: fsm.Transition{
			From: StateAgreed, Event: EventGoodsRefund, To: StateWaitDeliver, Action: action.GoodsRefund, Processor: nil,
		},
		// 取消事件：已通过 ---> 已取消
		EventCancel: fsm.Transition{
			From: StateAgreed, Event: EventCancel, To: StateCancel, Action: action.Cancel, Processor: nil,
		},
	},
	StateWaitDeliver: {
		// 发货事件：退货中 ---> 待收货
		EventDeliver: fsm.Transition{
			From: StateWaitDeliver, Event: EventDeliver, To: StateWaitReceive, Action: action.Delive,
			Processor: &processor.DeliveEventProcessor{},
		},
		// 取消事件：退货中 ---> 已取消
		EventCancel: fsm.Transition{
			From: StateWaitDeliver, Event: EventCancel, To: StateCancel, Action: action.Cancel, Processor: nil,
		},
	},
	StateWaitReceive: {
		// 签收事件：待收货 ---> 退款中
		EventSigned: fsm.Transition{
			From: StateWaitReceive, Event: EventSigned, To: StateWaitRefund, Action: action.Signed, Processor: nil,
		},
	},
	StateWaitRefund: {
		// 退款完成事件：退款中 ---> 已完成
		EventRefundCompleted: fsm.Transition{
			From: StateWaitRefund, Event: EventRefundCompleted, To: StateCompleted, Action: action.RefundCompleted,
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
	sm.SetName("退款状态图表")
	sm.SetStart(StateWaitApprove)
	sm.SetEnd(StateCompleted)
	sm.SetStates(StateDesc)
	sm.SetTransitions(transitions)
	return sm
}
