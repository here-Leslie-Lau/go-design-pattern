package builder

import (
	"fmt"
	"testing"
)

// 测试儿童套餐
func TestKids(t *testing.T) {
	director := director{}
	builder := newKidsComboBuilder()
	director.construct(builder)
	result := builder.getResult()
	if result.hamburger != 1 || result.coke != 1 {
		fmt.Println("这个小朋友吃太多了！")
		t.Fail()
	}
}

// 测试成人套餐
func TestAdult(t *testing.T) {
	director := director{}
	builder := newAdultComboBuilder()
	director.construct(builder)
	result := builder.getResult()
	if result.hamburger != 3 || result.coke != 3 {
		fmt.Println("这个成人吃太少了！")
		t.Fail()
	}
}
