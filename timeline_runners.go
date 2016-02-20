package main

import "sync"

/* timeline_runners.go contains a type that maintains a list of TimelineRunners and
   coordinates safe addition and removal of elements */

type TR struct {
	d string
	t string
}

type TimelineRunners struct {
	runners map[TR]*TimelineRunner
	mutex   sync.Mutex
}

func (tr *TimelineRunners) AddOrReplaceTimeline(tl Timeline) {
	(&tr.mutex).Lock()
	defer (&tr.mutex).Unlock()

	d := tl.DinoName
	t := tl.Trigger

	// lazily instantiate the map
	if tr.runners == nil {
		tr.runners = make(map[TR]*TimelineRunner)
	}

	// do we have an existing runner for this d, t pair?
	pair := TR{d, t}
	runner, ok := tr.runners[pair]

	if ok {
		// Yes, so kill it
		runner.Die()
		// And remove it from the map
		delete(tr.runners, pair)
	}

	// create a new runner for this timeline
	runner = NewTimelineRunner(tl)
	runner.Live()
	tr.runners[pair] = runner
}

func (tr *TimelineRunners) DeleteAllTimelines() {
	(&tr.mutex).Lock()
	defer (&tr.mutex).Unlock()

	// lazily instantiate the map
	if tr.runners == nil {
		tr.runners = make(map[TR]*TimelineRunner)
	}

	for pair, runner := range tr.runners {
		runner.Die()
		delete(tr.runners, pair)
	}
}

func (tr *TimelineRunners) GetAllTimelines() []Timeline {
	// We want to *copy* the data
	(&tr.mutex).Lock()
	defer (&tr.mutex).Unlock()

	// lazily instantiate the map
	if tr.runners == nil {
		tr.runners = make(map[TR]*TimelineRunner)
	}

	list := make([]Timeline, 0, len(tr.runners))
	for _, runner := range tr.runners {
		list = append(list, runner.T)
	}

	return list
}
