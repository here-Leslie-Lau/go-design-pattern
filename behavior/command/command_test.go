package command

import "testing"

func TestOrder(t *testing.T) {
	// construct
	d := &dish{names: []string{"红烧肉", "娃娃菜", "福建人???"}}
	// 下单命令
	o := &orderCommand{d: d}
	// 将下单命令交给服务员
	w := waiter{}
	w.setCommand(o)
	// 输出
	// waiter is actioning...
	// order command is starting...
	// making [红烧肉 娃娃菜 福建人???] now
	// order command is ending...
	// waiter is finished...
	if err := w.action(); err != nil {
		t.Fail()
	}
}
