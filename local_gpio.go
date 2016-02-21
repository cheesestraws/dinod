package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/kidoman/embd"
)

const debounceInterval = 20 * time.Millisecond

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

func (g *LocalGPIO) SetupInput(pin int, dinoName string, sensorName string) {
	(&g.mutex).Lock()
	defer (&g.mutex).Unlock()

	pname := fmt.Sprintf("P1_%d", pin)
	p, err := embd.NewDigitalPin(pname)
	if err != nil {
		fmt.Printf("Could not set %v to output: %v", pname, err)
	}
	p.SetDirection(embd.In)
	p.PullUp()

	// debouncing
	lastTrigger := time.Now()

	// Watch interrupts are all triggered on one goroutine so we're fairly safe to
	// just use Watch()
	p.Watch(embd.EdgeRising, func(p embd.DigitalPin) {
		if time.Since(lastTrigger) < debounceInterval {
			return
		}

		lastTrigger = time.Now()

		fmt.Printf("Sensor pin %v triggered.\n", pin)
		trigger(dinoName, sensorName)

	})

	fmt.Printf("Pin %v is now an input (prodding sensor %v / %v)\n", pname, dinoName, sensorName)

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
