package main

import "github.com/kidoman/embd" // ew this shouldn't be here TODO
import "encoding/json"
import "io/ioutil"

type State struct {
	timelines TimelineRunners
	dinos     Dinos
	gpio      GPIO
}

func (s *State) LoadConfig(filename string) error {
	// we should have a specific Config type, really
	str, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	var dinos Dinos
	err = json.Unmarshal([]byte(str), &dinos)
	if err != nil {
		return err
	}

	s.dinos = dinos
	return nil
}

func (s *State) Init() {
	host, _, _ := embd.DetectHost()
	if host == embd.HostNull {
		println("No real GPIO, using fake_gpio")
		s.gpio = GPIO(FakeGPIO{})
	} else {
		println("Using embd GPIO")
		s.gpio = GPIO(&LocalGPIO{})
	}

	SetupGPIOForDinos(s.dinos, s.gpio)
}

var state State

func getCurrentDinoList() Dinos {
	return state.dinos
}

func getCurrentTimelines() []Timeline {
	return state.timelines.GetAllTimelines()
}

func replaceAllTimelines(ts []Timeline) []Timeline {
	return []Timeline{}
}

func addOrReplaceTimeline(ts Timeline) []Timeline {
	state.timelines.AddOrReplaceTimeline(ts)
	return state.timelines.GetAllTimelines()
}

func deleteAllTimelines() []Timeline {
	state.timelines.DeleteAllTimelines()
	return state.timelines.GetAllTimelines()
}

func trigger(dino, sensor string) {
	state.timelines.Trigger(dino, sensor)
}
