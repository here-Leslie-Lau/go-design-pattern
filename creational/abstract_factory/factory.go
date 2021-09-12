package abstract_factory

// 工厂的抽象，提供两个方法(生成手机与手表)
type factory interface {
	getWatch() iWatch
	getPhone() phone
}

// appleFactory 苹果产品工厂
type appleFactory struct {}

func (f *appleFactory) getWatch() iWatch {
	return &AppleWatch{watch: watch{name: "apple"}}
}

func (f *appleFactory) getPhone() phone {
	return &iphone{mobile: mobile{name: "apple"}}
}



// xiaoMiFactory 小米产品工厂
type xiaoMiFactory struct {}

func (f *xiaoMiFactory) getWatch() iWatch {
	return &MiWatch{watch: watch{name: "xiaoMi"}}
}

func (f *xiaoMiFactory) getPhone() phone {
	return &xiaoMi{mobile: mobile{name: "xiaoMi"}}
}
