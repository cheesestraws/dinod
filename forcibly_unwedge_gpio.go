package main

// This is *awful*.

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

var gpioRE = regexp.MustCompile(`^gpio(\d+)$`)

func unwedgePin(pin string) error {
	fh, err := os.OpenFile("/sys/class/gpio/unexport", os.O_WRONLY, os.ModeExclusive)
	if err != nil {
		return err
	}
	defer fh.Close()
	_, err = fh.WriteString(pin)
	return err
}

func forciblyUnwedgeGPIO() {
	println("Forcibly unwedging GPIO")

	info, err := ioutil.ReadDir("/sys/class/gpio")
	if err != nil {
		fmt.Printf("Unwedging GPIO failed: %v.  Proceeding in hope anyway.\n", err)
		return
	}

	for _, i := range info {
		matches := gpioRE.FindStringSubmatch(i.Name())
		if len(matches) > 1 {
			fmt.Printf("Unwedging %v\n", matches[1])
			err := unwedgePin(matches[1])
			if err != nil {
				fmt.Printf("Unwedging pin %v failed: %v.\n", matches[1], err)
			}
		}
	}

	println("Dinod *might* work now.")
}
