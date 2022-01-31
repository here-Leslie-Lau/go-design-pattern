package flyweight

// dress 享元接口
type dress interface {
	getColor() string
}

// redTeamDress 红队服装享元接口实现
type redTeamDress struct{}

func (d *redTeamDress) getColor() string {
	return "red"
}

// blueTeamDress 蓝队服装享元接口实现
type blueTeamDress struct{}

func (d *blueTeamDress) getColor() string {
	return "blue"
}

// newDressFactory 初始化享元工厂对象
func newDressFactory() map[string]dress {
	m := map[string]dress{"red": &redTeamDress{}, "blue": &blueTeamDress{}}
	return m
}

// person 队员
type person struct {
	dress dress
}
