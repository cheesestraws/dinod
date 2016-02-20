package main

import "time"

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
			break
		}
	}

	tr.DeadC <- struct{}{}
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

	print("tick\n")

	// do magical shit here

	tr.Step += 1
	if tr.Step == tr.T.Length {
		tr.Running = false
		tr.Step = 0
	}

}
