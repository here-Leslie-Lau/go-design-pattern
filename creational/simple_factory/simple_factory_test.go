package simple_factory

import "testing"

func TestBMW(t *testing.T) {
	car := NewCar("bmw")
	if car == nil {
		t.Fail()
	}
	car.open()
	// 运行结果:bmw启动
}

func TestBenZ(t *testing.T) {
	car := NewCar("benz")
	if car == nil {
		t.Fail()
	}
	car.open()
	// 运行结果:benz启动
}
