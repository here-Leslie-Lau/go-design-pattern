package simple_factory

import "fmt"

type car interface {
	open()
}

// BMW 实现Car
type BMW struct {
	name string
}

func NewCar(name string) car {
	switch name {
	case "bmw":
		return &BMW{name: name}
	case "benz":
		return &BenZ{name: name}
	}
	return nil
}

// Open bmw重写Open
func (b *BMW) open() {
	fmt.Println(b.name + "启动")
}

// BenZ 实现Car
type BenZ struct {
	name string
}

// Open benz重写Open
func (b *BenZ) open() {
	fmt.Println(b.name + "启动")
}
