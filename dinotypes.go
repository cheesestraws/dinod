package main

import (
	"errors"
	"fmt"
)

type Dino struct {
	Name         string         `json:"name"`
	FriendlyName string         `json:"friendlyName"`
	ImageURI     string         `json:"imageURI"`
	Sensors      []DinoSensor   `json:"sensors"`
	Actuators    []DinoActuator `json:"actuators"`
}

func (d Dino) Valid() error {
	if d.Name == "" {
		return errors.New("dinosaur has no name")
	}

	for _, sensor := range d.Sensors {
		err := sensor.Valid()
		if err != nil {
			return err
		}
	}

	for _, actuator := range d.Actuators {
		err := actuator.Valid()
		if err != nil {
			return err
		}
	}

	return nil
}

type Dinos []Dino

func (d Dinos) FindDino(name string) *Dino {
	for _, dino := range d {
		if dino.Name == name {
			return &dino
		}
	}
	return nil
}

type DinoSensor struct {
	Name         string `json:"name"`
	FriendlyName string `json:"friendlyName"`
	Type         string `json:"type"`
	Pin          int    `json:"pin"`
}

func (d DinoSensor) Valid() error {
	if d.Name == "" {
		return errors.New("sensor has no name")
	}

	if d.Type != "pulse" {
		return errors.New(fmt.Sprintf("sensor %v has an invalid type", d.Name))
	}

	return nil
}

type DinoActuator struct {
	Name         string `json:"name"`
	FriendlyName string `json:"friendlyName"`
	Type         string `json:"type"`
	Pin          int    `json:"pin"`
}

func (d DinoActuator) Valid() error {
	if d.Name == "" {
		return errors.New("actuator has no name")
	}

	if d.Type != "onoff" {
		return errors.New(fmt.Sprintf("actuator %v has an invalid type", d.Name))
	}

	return nil
}
