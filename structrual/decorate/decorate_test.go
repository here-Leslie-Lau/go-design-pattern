package decorate

import (
	"testing"
)

func TestDecorate(t *testing.T) {
	var goods goods
	// 这里我们没有直接用cola或drink,而是用装饰过的decorate,减少客户端耦合
	goods = newDecorator()
	price := goods.GetPrice()

	if price != 3 {
		t.Fail()
	}
}
