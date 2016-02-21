package main

import "github.com/kidoman/embd" // ew this shouldn't be here TODO

type State struct {
	timelines TimelineRunners
	dinos     Dinos
	gpio      GPIO
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

var state State = State{
	dinos: []Dino{
		Dino{
			Name:         "d",
			FriendlyName: "test dino",
			Sensors: []DinoSensor{
				DinoSensor{
					Name:         "s",
					FriendlyName: "sensor",
					Type:         "none",
					Pin:          3,
				},
			},
			Actuators: []DinoActuator{
				DinoActuator{
					Name:         "red",
					FriendlyName: "actuator",
					Type:         "none",
					Pin:          33,
				},
				DinoActuator{
					Name:         "amber",
					FriendlyName: "actuator",
					Type:         "none",
					Pin:          35,
				},
				DinoActuator{
					Name:         "green",
					FriendlyName: "actuator",
					Type:         "none",
					Pin:          37,
				},
			},
		},
	},
}

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
