package abstract_factory

import (
	"fmt"
	"time"
)

// iWatch 手表产品的接口
type iWatch interface {
	watchTime()
}

type watch struct{
	name string
}

func (w *watch) watchTime() {
	fmt.Printf("name:%s,time:%s\n", w.name, time.Now().Format("2006-01-02"))
}

// AppleWatch 苹果手表产品的实现
type AppleWatch struct {
	watch
}

// MiWatch 小米手表产品的实现
type MiWatch struct {
	watch
}

// phone 手机产品的接口
type phone interface {
	playGame()
}

type mobile struct{
	name string
}

func (m *mobile) playGame() {
	fmt.Println("play game,name:", m.name)
}

// 苹果手机产品的实现
type iphone struct {
	mobile
}

// 小米手机产品的实现
type xiaoMi struct {
	mobile
}
