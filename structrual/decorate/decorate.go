package decorate

// goods component部件
type goods interface {
	GetPrice() float64
}

// cola concreteComponent - cola
type cola struct {
	price float64
}

func (c *cola) SetPrice(price float64) {
	c.price = price
}

func (c *cola) GetPrice() float64 {
	return c.price
}

// drink concreteComponent - drink
type drink struct {
	price float64
}

func (d *drink) SetPrice(price float64) {
	d.price = price
}

func (d *drink) GetPrice() float64 {
	return d.price
}

// Decorator decorate - 实现了component部件接口,可以封装额外功能
type Decorator struct {
	cola  goods
	drink goods
}

func newDecorator() *Decorator {
	return &Decorator{cola: &cola{price: 1}, drink: &drink{price: 5}}
}

func (d *Decorator) GetPrice() float64 {
	// 可乐与酒一块购买可以打5折
	result := (d.cola.GetPrice() + d.drink.GetPrice()) * 0.5
	return result
}
