package action

import (
	fsm "github.com/WTongStudio/fsm-order-demo"
	"github.com/WTongStudio/fsm-order-demo/helper"
)

// Agree 通过
func Agree(from fsm.State, event fsm.Event, to fsm.State) error {
	helper.Log("通过申请，旧状态:%d，事件: %s，新状态: %d", from, event, to)
	return nil
}
