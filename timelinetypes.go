package main

import "fmt"

type Timeline struct {
	FriendlyName string           `json:"friendlyName"`
	DinoName     string           `json:"dinoName"`
	Trigger      string           `json:"trigger"`
	Length       int              `json:"length"`
	TimePerStep  float64          `json:"timePerStep"`
	Timeline     map[string][]int `json:"timeline"`
}

func (t Timeline) ValidateAgainst(dinos Dinos) error {
	if t.DinoName == "" {
		return fmt.Errorf("no dinoName provided")
	}

	dino := dinos.FindDino(t.DinoName)
	if dino == nil {
		return fmt.Errorf("dinoName %v supplied but this dino does not exist", t.DinoName)
	}

	// check if trigger exists
	foundTrigger := false
	for _, sensor := range dino.Sensors {
		if sensor.Name == t.Trigger {
			foundTrigger = true
		}
	}

	if !foundTrigger {
		return fmt.Errorf("Trigger sensor %v does not exist in dino %v", t.Trigger, t.DinoName)
	}

	// Check that each actuator exists
	// make a set of the actuators on the dino
	actuators := make(map[string]struct{})
	for _, actuator := range dino.Actuators {
		actuators[actuator.Name] = struct{}{}
	}

	// and check each actuator on the timeline matches
	for actuator, _ := range t.Timeline {
		_, ok := actuators[actuator]
		if !ok {
			return fmt.Errorf("Timeline refers to actuator %v which is not present in dino %v", actuator, t.DinoName)
		}
	}

	return nil
}

func (t Timeline) Valid() error {
	return t.ValidateAgainst(getCurrentDinoList())
}
