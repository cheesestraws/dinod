package main

/*
fake_gpio.go contains a stub implementation of the GPIO interface (see gpio.go).  This
implementation does nothing except for dump state changes to stdout.  This implementation
is used for debugging directly-connected Dinos on a laptop.
*/

import (
	"fmt"
)

/*
The FakeGPIO type contains the state necessary to maintain the fake GPIO implementation:
that is to say, none at all.  Instances of FakeGPIO implement the GPIO interface.
*/
type FakeGPIO struct{}

// compile time check it complies with gpio interface
var _ = GPIO(FakeGPIO{})

/*
Init initialises a fake GPIO instance.  It does basically nothing.
*/
func (g FakeGPIO) Init() {
	println("Initialising FakeGPIO...")
}

/*
SetupInput sets up a fake GPIO input pin.  It does basically nothing.
*/
func (g FakeGPIO) SetupInput(pin int, dinoName string, sensorName string) {
	fmt.Printf("Pin %v is henceforth an input (%v, %v)\n", pin, dinoName, sensorName)
}

/*
SetupOutput sets up a fake GPIO output pin.  It does basically nothing.
*/
func (g FakeGPIO) SetupOutput(pin int) {
	fmt.Printf("Pin %v is henceforth an output\n", pin)
}

/*
SetPin sets the value of a fake GPIO output pin.  It does basically nothing.
*/
func (g FakeGPIO) SetPin(pin int, value int) {
	fmt.Printf("Pin %v => %v\n", pin, value)
}
