package chain_responsibility

import "testing"

func TestChainResponsibility(t *testing.T) {
	// construct
	doctor := &doctor{}
	reception := &reception{}
	reception.setNext(doctor)
	p := patient{}

	reception.execute(&p)
	// output:
	// reception is handing, patient: &{}
	// doctor is handing, patient: &{}
}
