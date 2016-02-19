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

func (t Timeline) Validate(dinos Dinos) error {
	if t.DinoName == "" {
		return fmt.Errorf("no dinoName provided")
	}

	if dinos.FindDino(t.DinoName) == nil {
		return fmt.Errorf("dinoName %v supplied but this dino does not exist", t.DinoName)
	}

	return nil
}
