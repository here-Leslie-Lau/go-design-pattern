package prototype

// cloneable 原型接口
type cloneable interface {
	clone() cloneable
}

type car struct {
	name string
}

func newCar(name string) *car {
	return &car{name: name}
}

func (c *car) clone() cloneable {
	copy := *c
	return &copy
}
