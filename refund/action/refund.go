package action

import (
	fsm "github.com/WTongStudio/fsm-order-demo"
	"github.com/WTongStudio/fsm-order-demo/helper"
)

// Refund 退款
func Refund(from fsm.State, event fsm.Event, to fsm.State) error {
	helper.Log("售后退款，旧状态:%d，事件: %s，新状态: %d", from, event, to)
	return nil
}
