package abstract_factory

import (
	"testing"
)

func TestApple(t *testing.T) {
	var factory factory
	factory = &appleFactory{}
	phone := factory.getPhone()
	watch := factory.getWatch()
	// 运行结果: play game,name: apple
	phone.playGame()
	// 运行结果: name:apple,time:2021-09-12
	watch.watchTime()
}

func TestXiaoMi(t *testing.T) {
	var factory factory
	factory = &xiaoMiFactory{}
	phone := factory.getPhone()
	watch := factory.getWatch()
	// 运行结果: play game,name: xiaoMi
	phone.playGame()
	// 运行结果: name:xiaoMi,time:2021-09-12
	watch.watchTime()
}
