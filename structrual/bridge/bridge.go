package bridge

import "fmt"

// game 游戏 - 实现部分
type game interface {
	playGame()
}

// nba2K 具体实现
type nba2K struct{}

func (n *nba2K) playGame() {
	fmt.Println("play game:nba2K")
}

// gta5 具体实现
type gta5 struct{}

func (g *gta5) playGame() {
	fmt.Println("play game:gta5")
}

// gameMachine 游戏机 - 抽象部分
type gameMachine interface {
	play()
	loadGame(game game) // 加载游戏
}

// xBox xBox游戏机 - 精确实现
type xBox struct {
	game game
}

func newXBox() *xBox {
	return &xBox{}
}

func (x *xBox) play() {
	fmt.Println("i am xBox")
	x.game.playGame()
}

func (x *xBox) loadGame(game game) {
	x.game = game
}

// ps5 ps5游戏机 - 精确实现
type ps5 struct {
	game game
}

func (p *ps5) play() {
	fmt.Println("i am ps5")
	p.game.playGame()
}

func (p *ps5) loadGame(game game) {
	p.game = game
}
