package chain_responsibility

import "fmt"

// patient 病人 - 请求处理对象
type patient struct{}

// department 责任链通用接口
type department interface {
	execute(p *patient)
	setNext(d department)
}

type reception struct {
	next department
}

func (r *reception) execute(p *patient) {
	fmt.Println("reception is handing, patient:", p)
	r.next.execute(p)
}

func (r *reception) setNext(d department) {
	r.next = d
}

type doctor struct{}

func (d *doctor) execute(p *patient) {
	fmt.Println("doctor is handing, patient:", p)
}

func (d *doctor) setNext(_ department) {
	panic("implement me")
}
