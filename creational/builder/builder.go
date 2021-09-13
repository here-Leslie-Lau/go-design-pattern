package builder

// builder接口,提供一个获取产品的方法与若干设置属性的方法
type iBuilder interface {
	setHamburger()
	setCoke()
	getResult() *combo
}

// combo - 套餐产品,其中包含可乐与汉堡
type combo struct {
	coke, hamburger uint
}

// comboBuilder builder的具体实现 - 儿童套餐
type kidsComboBuilder struct {
	c *combo
}

func newKidsComboBuilder() *kidsComboBuilder {
	return &kidsComboBuilder{c: &combo{}}
}

// 儿童套餐,一个汉堡就够了,可乐也同理
func (c *kidsComboBuilder) setHamburger() {
	c.c.hamburger = 1
}

func (c *kidsComboBuilder) setCoke() {
	c.c.coke = 1
}

func (c *kidsComboBuilder) getResult() *combo {
	return c.c
}

// comboBuilder builder的具体实现 - 成人套餐
type adultComboBuilder struct {
	c *combo
}

func newAdultComboBuilder() *adultComboBuilder {
	return &adultComboBuilder{c: &combo{}}
}

// 成人套餐,一个汉堡不够,要三个,可乐也是
func (c *adultComboBuilder) setHamburger() {
	c.c.hamburger = 3
}

func (c *adultComboBuilder) setCoke() {
	c.c.coke = 3
}

func (c *adultComboBuilder) getResult() *combo {
	return c.c
}
