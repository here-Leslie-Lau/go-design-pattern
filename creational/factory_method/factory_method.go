package factory_method

import "fmt"

type car interface {
	open()
	setName(name string)
}

// BMW car的实现类
type BMW struct {
	name string
}

func (b *BMW) open() {
	fmt.Println("启动了:", b.name)
}

func (b *BMW) setName(name string) {
	b.name = name
}

// BenZ car的实现类
type BenZ struct {
	name string
}

func (b *BenZ) open() {
	fmt.Println("启动了:", b.name)
}

func (b *BenZ) setName(name string) {
	b.name = name
}
