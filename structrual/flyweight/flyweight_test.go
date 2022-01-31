package flyweight

import (
	"fmt"
	"testing"
)

func TestFlyWeight(t *testing.T) {
	// 获取享元对象工厂
	f := newDressFactory()
	// 获取红队成员服装
	red := &person{dress: f["red"]}
	fmt.Println(red.dress.getColor())
	// 获取蓝队成员服装
	blue := person{dress: f["blue"]}
	fmt.Println(blue.dress.getColor())
	// 生成第二名蓝队成员，并比较服装对象是否仍然是同一个
	blue2 := &person{dress: f["blue"]}
	if blue2.dress != blue.dress {
		t.Fail()
	}
}
