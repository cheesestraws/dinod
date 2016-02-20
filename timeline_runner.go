package main

import "time"
import "fmt"

type TimelineRunner struct {
	T       Timeline
	StopC   chan struct{}
	StartC  chan struct{}
	DieC    chan struct{}
	DeadC   chan struct{}
	Step    int
	Running bool
	Alive   bool
}

func NewTimelineRunner(t Timeline) *TimelineRunner {
	return &TimelineRunner{
		T:       t,
		StopC:   make(chan struct{}),
		StartC:  make(chan struct{}),
		DieC:    make(chan struct{}),
		DeadC:   make(chan struct{}),
		Step:    0,
		Running: false,
		Alive:   false,
	}
}

func (tr *TimelineRunner) Live() {
	tr.Alive = true
	go tr.runLoop()
}

func (tr *TimelineRunner) Die() {
	tr.DieC <- struct{}{}
	// Wait for the death to happen
	<-tr.DeadC
}

func (tr *TimelineRunner) Start() {
	tr.StartC <- struct{}{}
}

func (tr *TimelineRunner) Stop() {
	tr.StopC <- struct{}{}
}

// run this only as a goroutine
func (tr *TimelineRunner) runLoop() {
	ticker := time.NewTicker(time.Duration(tr.T.TimePerStep * float64(time.Second)))
	for {
		select {
		case <-ticker.C:
			tr.runStep()
		case <-tr.StartC:
			tr.Running = true
		case <-tr.StopC:
			tr.Running = false
			tr.Step = 0
		case <-tr.DieC:
			tr.Running = false
			tr.Alive = false
			ticker.Stop()
			tr.DeadC <- struct{}{}
		}
	}
}

// run this only from the runLoop goroutine
func (tr *TimelineRunner) runStep() {
	if !tr.Running {
		return
	}
	// guard at the start in case something awful has happened
	if tr.Step == tr.T.Length {
		tr.Running = false
		tr.Step = 0
		return
	}

	for actuatorName, actuatorSlice := range tr.T.Timeline {
		var value int

		if len(actuatorSlice) < tr.Step {
			value = actuatorSlice[len(actuatorSlice)-1]
		} else {
			value = actuatorSlice[tr.Step]
		}

		// find the actuator pin
		dino := state.dinos.FindDino(tr.T.DinoName)
		if dino == nil {
			println("Dino is nil: something damned strange has happened.")
			break
		}

		actuator := dino.FindActuator(actuatorName)
		if actuator == nil {
			println("Actuator is nil: something damned strange has happened.")
			break
		}

		pin := actuator.Pin

		fmt.Printf("(%v) %v ~ %v => %v\n", tr.T.DinoName, pin, actuatorName, value)

		state.gpio.SetPin(pin, value)
	}

	tr.Step += 1
	if tr.Step == tr.T.Length {
		tr.Running = false
		tr.Step = 0
	}

}
