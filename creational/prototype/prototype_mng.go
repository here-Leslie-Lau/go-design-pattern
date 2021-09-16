package prototype

// prototypeManager 原型接口管理者,提供set&get
type prototypeManager struct {
	m map[string]cloneable
}

func newPrototypeManager() *prototypeManager {
	return &prototypeManager{m: make(map[string]cloneable)}
}

func (p *prototypeManager) get(name string) cloneable {
	return p.m[name].clone()
}

func (p *prototypeManager) set(name string, obj cloneable) {
	p.m[name] = obj
}
