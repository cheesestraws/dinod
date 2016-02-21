package main

import (
	"fmt"
	"sync"

	"github.com/kidoman/embd"
)

type PI struct { // a pair of integers
	a int
	b int
}

type LocalGPIO struct {
	mutex sync.Mutex
}

// compile time check it complies with gpio interface
var _ = GPIO(&LocalGPIO{})

func (g *LocalGPIO) Init() {
	(&g.mutex).Lock()
	defer (&g.mutex).Unlock()
	embd.InitGPIO()
}

func (g *LocalGPIO) SetupInput(pin int) {
	(&g.mutex).Lock()
	defer (&g.mutex).Unlock()

	pname := fmt.Sprintf("P1_%d", pin)
	p, err := embd.NewDigitalPin(pname)
	if err != nil {
		fmt.Printf("Could not set %v to output: %v", pname, err)
	}
	p.SetDirection(embd.In)
	p.PullUp()

	// Watch interrupts are all triggered on one goroutine so we're fairly safe to
	// just use Watch()
	p.Watch(embd.EdgeRising, func(p embd.DigitalPin) {
		fmt.Printf("Sensor pin %v triggered.\n", pin)
	})

	fmt.Printf("Pin %v is now an input\n", pname)

}

func (g *LocalGPIO) SetupOutput(pin int) {
	(&g.mutex).Lock()
	defer (&g.mutex).Unlock()
	pname := fmt.Sprintf("P1_%d", pin)
	p, err := embd.NewDigitalPin(pname)
	if err != nil {
		fmt.Printf("Could not set %v to output: %v", pname, err)
	}
	p.SetDirection(embd.Out)
	fmt.Printf("Pin %v is now an output\n", pname)
}

func (g *LocalGPIO) SetPin(pin int, value int) {
	(&g.mutex).Lock()
	defer (&g.mutex).Unlock()
	pname := fmt.Sprintf("P1_%d", pin)
	var pushValue int
	if value <= 0 {
		pushValue = embd.Low
	} else {
		pushValue = embd.High
	}
	embd.DigitalWrite(pname, pushValue)
}
