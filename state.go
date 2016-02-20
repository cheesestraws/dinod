package main

type State struct {
	timelines TimelineRunners
	dinos     []Dino
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
					Pin:          1,
				},
			},
			Actuators: []DinoActuator{
				DinoActuator{
					Name:         "a",
					FriendlyName: "actuator",
					Type:         "none",
					Pin:          2,
				},
				DinoActuator{
					Name:         "b",
					FriendlyName: "actuator",
					Type:         "none",
					Pin:          2,
				},
				DinoActuator{
					Name:         "c",
					FriendlyName: "actuator",
					Type:         "none",
					Pin:          2,
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
