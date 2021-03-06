package main

import "github.com/kidoman/embd" // ew this shouldn't be here TODO
import "encoding/json"
import "io/ioutil"

type State struct {
	timelines    TimelineRunners
	timelineFile string
	dinos        DinoStates
	rawDinos     Dinos
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

	s.rawDinos = dinos

	for _, d := range dinos {
		ds := DinoState{d, nil}
		s.dinos = append(s.dinos, ds)
	}

	return nil
}

func (s *State) RestoreTimelines(filename string) error {
	s.timelineFile = filename

	str, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	var tls []Timeline
	err = json.Unmarshal([]byte(str), &tls)
	if err != nil {
		return err
	}

	state.timelines.ReplaceAllTimelines(tls)
	return nil
}

func (s *State) SaveTimelines() error {
	tls := s.timelines.GetAllTimelines()
	json, _ := json.Marshal(tls)
	return ioutil.WriteFile(s.timelineFile, json, 0664)
}

func (s *State) Init() {
	for i := range s.dinos {
		gpio := s.gpioBackendFor(s.dinos[i].Dino)
		SetupGPIOForDino(s.dinos[i].Dino, gpio)
		s.dinos[i].gpio = gpio
	}
}

func (s *State) gpioBackendFor(d Dino) GPIO {
	var gpio GPIO

	if d.Backend == "mqtt" {
		gpio = GPIO(&MQTT_GPIO{})
	} else {
		host, _, _ := embd.DetectHost()
		if host == embd.HostNull {
			println("No real GPIO, using fake_gpio")
			gpio = GPIO(FakeGPIO{})
		} else {
			println("Using embd GPIO")
			gpio = GPIO(&LocalGPIO{})
		}
	}

	gpio.Init()

	return gpio
}

var state State

func getCurrentDinoList() Dinos {
	return state.rawDinos
}

func getCurrentTimelines() []Timeline {
	return state.timelines.GetAllTimelines()
}

func replaceAllTimelines(ts []Timeline) []Timeline {
	state.timelines.ReplaceAllTimelines(ts)
	return state.timelines.GetAllTimelines()
}

func addOrReplaceTimeline(ts Timeline) []Timeline {
	state.timelines.AddOrReplaceTimeline(ts)
	state.SaveTimelines()
	return state.timelines.GetAllTimelines()
}

func deleteAllTimelines() []Timeline {
	state.timelines.DeleteAllTimelines()
	return state.timelines.GetAllTimelines()
}

func trigger(dino, sensor string) {
	state.timelines.Trigger(dino, sensor)
}
