package builder

type director struct {}

func newDirector() *director {
	return &director{}
}

// director,这个做成接口会比较好,但是为了不复杂,就简化了
func (d *director) construct(builder iBuilder) {
	builder.setCoke()
	builder.setHamburger()
}
