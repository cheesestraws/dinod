package main

import (
	"fmt"
)

type FakeGPIO struct{}

// compile time check it complies with gpio interface
var _ = GPIO(FakeGPIO{})

func (g FakeGPIO) Init() {
	println("Initialising FakeGPIO...")
}

func (g FakeGPIO) SetupInput(pin int, dinoName string, sensorName string) {
	// do nothing for the moment
	fmt.Printf("Pin %v is henceforth an output (%v, %v)\n", pin, dinoName, sensorName)
}

func (g FakeGPIO) SetupOutput(pin int) {
	fmt.Printf("Pin %v is henceforth an output\n", pin)
}

func (g FakeGPIO) SetPin(pin int, value int) {
	fmt.Printf("Pin %v => %v\n", pin, value)
}
