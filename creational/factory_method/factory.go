package factory_method

// 这个工厂接口用于创建产品
type factory interface {
	create() car
}

type bmwFactory struct{}

func (b *bmwFactory) create() car {
	car := &BMW{}
	car.setName("bmw")
	return car
}

type benzFactory struct{}

func (b *benzFactory) create() car {
	car := &BenZ{}
	car.setName("benz")
	return car
}
