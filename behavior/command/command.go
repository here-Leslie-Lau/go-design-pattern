package command

import "fmt"

// iCommand 命令对象
type iCommand interface {
	// 命令对象执行方法
	execute() error
}

// waiter 服务员
type waiter struct {
	c iCommand
}

func (w *waiter) setCommand(c iCommand) {
	w.c = c
}

func (w *waiter) action() error {
	if w.c == nil {
		return fmt.Errorf("waiter's command is empty")
	}
	fmt.Println("waiter is actioning...")
	if err := w.c.execute(); err != nil {
		return err
	}
	fmt.Println("waiter is finished...")
	return nil
}

// orderCommand 下单命令对象
type orderCommand struct {
	d *dish
}

func (o *orderCommand) execute() error {
	fmt.Println("order command is starting...")
	if err := o.d.make(); err != nil {
		return err
	}
	fmt.Println("order command is ending...")
	return nil
}

// dish receiver接受者类 - 处理菜品
type dish struct {
	names []string
}

func (d *dish) make() error {
	fmt.Println(fmt.Sprintf("making %v now", d.names))
	return nil
}
