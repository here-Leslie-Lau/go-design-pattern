package factory_method

import "testing"

func TestBmw(t *testing.T) {
	var factory factory
	factory = &bmwFactory{}
	car := factory.create()
	car.open()
	// 输出结果: 启动了: bmw
}

func TestBenz(t *testing.T) {
	var factory factory
	factory = &benzFactory{}
	car := factory.create()
	car.open()
	// 输出结果: 启动了: benz
}
