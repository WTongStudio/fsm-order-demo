package action

import (
	fsm "github.com/WTongStudio/fsm-order-demo"
	"github.com/WTongStudio/fsm-order-demo/helper"
)

// CancelRefund 取消退款
func CancelRefund(from fsm.State, event fsm.Event, to fsm.State) error {
	helper.Log("子订单取消售后，旧状态：%d，事件：%s，新状态：%d", from, event, to)
	return nil
}
