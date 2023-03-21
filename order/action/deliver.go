package action

import (
	fsm "github.com/WTongStudio/fsm-order-demo"
	"github.com/WTongStudio/fsm-order-demo/helper"
)

// Delive 发货
func Delive(from fsm.State, event fsm.Event, to fsm.State) error {
	helper.Log("子订单发货，旧状态：%d，事件：%s，新状态：%d", from, event, to)
	return nil
}
