package action

import (
	fsm "github.com/WTongStudio/fsm-order-demo"
	"github.com/WTongStudio/fsm-order-demo/helper"
)

// PayConfirm 支付确认
func PayConfirm(from fsm.State, event fsm.Event, to fsm.State) error {
	helper.Log("主订单支付确认，旧状态：%d，事件：%s，新状态：%d", from, event, to)
	return nil
}
