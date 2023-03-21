package action

import (
	fsm "github.com/WTongStudio/fsm-order-demo"
	"github.com/WTongStudio/fsm-order-demo/helper"
)

// Cancel 取消订单
func Cancel(from fsm.State, event fsm.Event, to fsm.State) error {
	helper.Log("取消售后，旧状态:%d，事件: %s，新状态: %d", from, event, to)
	return nil
}
