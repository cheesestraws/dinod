package main

import (
	"fmt"

	"github.com/kidoman/embd"
)

type LocalGPIO struct{}

// compile time check it complies with gpio interface
var _ = GPIO(LocalGPIO{})

func (g LocalGPIO) Init() {
	embd.InitGPIO()
}

func (g LocalGPIO) SetupInput(pin int) {
	// do nothing for the moment
}

func (g LocalGPIO) SetupOutput(pin int) {
	pname := fmt.Sprintf("P1_%d", pin)
	embd.SetDirection(pname, embd.Out)
	embd.DigitalWrite(pname, embd.Low)
}

func (g LocalGPIO) SetPin(pin int, value int) {
	pname := fmt.Sprintf("P1_%d", pin)
	var pushValue int
	if value <= 0 {
		pushValue = embd.Low
	} else {
		pushValue = embd.High
	}
	embd.DigitalWrite(pname, pushValue)
}
