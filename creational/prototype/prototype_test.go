package prototype

import (
	"testing"
)

func TestClone(t *testing.T) {
	manager := newPrototypeManager()
	car := newCar("bmw")

	manager.set("bmw", car)
	// clone出一个宝马
	copyCar := manager.get(car.name)
	// 测试两个地址是否一致,以此判断是否完成clone
	if copyCar == car {
		t.Fail()
	}
}
