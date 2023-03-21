package action

import (
	fsm "github.com/WTongStudio/fsm-order-demo"
	"github.com/WTongStudio/fsm-order-demo/helper"
)

// GoodsRefund 退货退款
func GoodsRefund(from fsm.State, event fsm.Event, to fsm.State) error {
	helper.Log("售后退款，旧状态:%d，事件: %s，新状态: %d", from, event, to)
	return nil
}
