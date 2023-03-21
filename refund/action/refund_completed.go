package action

import (
	fsm "github.com/WTongStudio/fsm-order-demo"
	"github.com/WTongStudio/fsm-order-demo/helper"
)

// RefundCompleted 退款完成
func RefundCompleted(from fsm.State, event fsm.Event, to fsm.State) error {
	helper.Log("退款完成，旧状态:%d，事件: %s，新状态: %d", from, event, to)
	return nil
}
