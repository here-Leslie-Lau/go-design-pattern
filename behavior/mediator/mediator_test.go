package mediator

import (
	"fmt"
	"sync"
	"testing"
)

func TestMeditor(t *testing.T) {
	// constructor
	one := doctorOne{
		name:  "one",
		mutex: sync.Mutex{},
	}
	two := doctorOne{
		name:  "two",
		mutex: sync.Mutex{},
	}
	m := map[string]doctor{one.name: &one, two.name: &two}
	mediator := newMediatorImpl(m)

	one.m = mediator
	two.m = mediator

	if err := one.diagnose("first patient"); err != nil {
		fmt.Println("one err", err)
	}
	if err := two.diagnose("second patient"); err != nil {
		fmt.Println("two err", err)
	}
}
