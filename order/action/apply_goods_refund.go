package action

import (
	fsm "github.com/WTongStudio/fsm-order-demo"
	"github.com/WTongStudio/fsm-order-demo/helper"
)

// ApplyGoodsRefund 申请退货退款
func ApplyGoodsRefund(from fsm.State, event fsm.Event, to fsm.State) error {
	helper.Log("子订单申请退货退款，旧状态：%d，事件：%s，新状态：%d", from, event, to)
	return nil
}
