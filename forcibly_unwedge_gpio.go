package main

// This is *awful*.

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

var gpioRE = regexp.MustCompile(`^gpio(\d+)$`)

func forciblyUnwedgeGPIO() {
	println("Forcibly unwedging GPIO")

	info, err := ioutil.ReadDir("/sys/class/gpio")
	if err != nil {
		fmt.Printf("Unwedging GPIO failed: %v.  Proceeding in hope anyway.", err)
		return
	}

	for _, i := range info {
		matches := gpioRE.FindStringSubmatch(i.Name())
		if len(matches) > 1 {
			fmt.Printf("Unwedging %v", i.Name())
		}
	}

	println("Dinod *might* work now.")
}
