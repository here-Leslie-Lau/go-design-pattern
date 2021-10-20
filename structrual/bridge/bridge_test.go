package bridge

import "testing"

// 测试xBox调用
func TestXBox(t *testing.T) {
	var machine gameMachine
	var g game
	machine = &xBox{}
	// 先玩nba2k
	g = &nba2K{}
	machine.loadGame(g)
	// 输出:
	// i am xBox
	// play game:nba2K
	machine.play()
	// 再玩gta5
	g = &gta5{}
	machine.loadGame(g)
	// 输出:
	// i am xBox
	// play game:gta5
	machine.play()
}

// 测试ps5调用
func TestPs5(t *testing.T) {
	var machine gameMachine
	var g game
	machine = &ps5{}
	// 先玩nba2k
	g = &nba2K{}
	machine.loadGame(g)
	// 输出:
	// i am ps5
	// play game:nba2K
	machine.play()
	// 再玩gta5
	g = &gta5{}
	machine.loadGame(g)
	// 输出:
	// i am ps5
	// play game:gta5
	machine.play()
}
