package processor

import (
	fsm "github.com/WTongStudio/fsm-order-demo"
	"github.com/WTongStudio/fsm-order-demo/helper"
)

// EventProcessor 退款处理器
type EventProcessor struct{}

// ExitOldState 离开旧状态
func (p *EventProcessor) ExitOldState(state fsm.State, event fsm.Event) error {
	helper.Log("退款状态机默认处理器 -- 离开旧状态，状态: %d，事件: %s", state, event)
	return nil
}

// EnterNewState 进入新状态
func (p *EventProcessor) EnterNewState(state fsm.State, event fsm.Event) error {
	helper.Log("退款状态机默认处理器 -- 进入新状态，状态: %d，事件: %s", state, event)
	return nil
}
