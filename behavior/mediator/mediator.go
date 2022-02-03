package mediator

import (
	"fmt"
	"sync"
)

type doctor interface {
	// diagnose 诊断病人
	diagnose(patient string) error
	// setDiagnoseStatus 设置诊断状态
	setDiagnoseStatus(flag bool)
}

// mediator 中介者接口
type mediator interface {
	// notifyDiagnoseFinished 通知病人诊断完成
	notifyDiagnoseFinished(name string)
}

type doctorOne struct {
	// 是否在诊断
	isDiagnose bool
	name       string
	m          mediator
	mutex      sync.Mutex
}

func (d *doctorOne) diagnose(patient string) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	if d.isDiagnose {
		return fmt.Errorf("当前医生已被约满，请稍后再试")
	}
	fmt.Println("doctor one is diagnosing:", patient)
	d.setDiagnoseStatus(true)
	// notify
	d.m.notifyDiagnoseFinished(d.name)
	return nil
}

func (d *doctorOne) setDiagnoseStatus(flag bool) {
	d.isDiagnose = flag
}

type doctorTwo struct {
	// 是否在诊断
	isDiagnose bool
	name       string
	m          mediator
	mutex      sync.Mutex
}

func (d *doctorTwo) diagnose(patient string) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	if !d.isDiagnose {
		return fmt.Errorf("当前医生已被约满，请稍后再试")
	}
	fmt.Println("doctor two is diagnosing:", patient)
	d.setDiagnoseStatus(true)
	// notify
	d.m.notifyDiagnoseFinished(d.name)
	return nil
}

func (d *doctorTwo) setDiagnoseStatus(flag bool) {
	d.isDiagnose = flag
}

type mediatorImpl struct {
	m   map[string]doctor
}

func newMediatorImpl(m map[string]doctor) *mediatorImpl {
	return &mediatorImpl{m: m}
}

func (m *mediatorImpl) notifyDiagnoseFinished(name string) {
	m.m[name].setDiagnoseStatus(false)
	fmt.Println("doctor had finished to diagnose:", name)
}
