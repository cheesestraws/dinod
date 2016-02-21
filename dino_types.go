package main

/*
dino_types.go contains type definitions for dinosaurs (or, if you must, something less
interesting than dinosaurs)
*/

import (
	"errors"
	"fmt"
)

/*
A Dino struct contains the definition of a dinosaur (or, I guess, some other robot or
device.  It does not contain the *runtime* state necessary to run the device, but only
the static data that can exist between runs.  The 'backend' field is additional to the
fields defined in the Dinosaur API.  This field defines the protocol that the dinod
instance should use to talk to the device.
*/
type Dino struct {
	Name         string         `json:"name"`
	FriendlyName string         `json:"friendlyName"`
	ImageURI     string         `json:"imageURI"`
	Sensors      []DinoSensor   `json:"sensors"`
	Actuators    []DinoActuator `json:"actuators"`
	Backend      string         `json:"backend"`
}

/*
Valid checks whether a Dino is valid.  If it is valid, it returns nil: if it is not,
it returns an error describing the problem that can be passed back to the user.
*/
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

/*
FindSensor returns the sensor description that corresponds to a given name for this Dino,
or nil if none is found.
*/
func (d Dino) FindSensor(name string) *DinoSensor {
	for _, sensor := range d.Sensors {
		if sensor.Name == name {
			return &sensor
		}
	}

	return nil
}

/*
FindActuator returns the actuator description that corresponds to a given name for this
Dino, or nil if none is found.
*/
func (d Dino) FindActuator(name string) *DinoActuator {
	for _, actuator := range d.Actuators {
		if actuator.Name == name {
			return &actuator
		}
	}

	return nil
}

/*
The Dinos type contains a list of Dino definitions.  It exists so that it can have method
definitions upon it.
*/
type Dinos []Dino

/*
FindDino returns the Dino definition with the given name from a list of Dino definitions.
If no such dino exists, it returns nil.
*/
func (d Dinos) FindDino(name string) *Dino {
	for _, dino := range d {
		if dino.Name == name {
			return &dino
		}
	}
	return nil
}

/*
A DinoState struct contains a Dino and any state necessary to make it actually run.
*/
type DinoState struct {
	Dino
	gpio GPIO
}

/*
The DinoStates type contains a list of Dino runtime states.  It exists so that it can
have method definitions upon it.
*/
type DinoStates []DinoState

/*
FindDinoState returns the Dino state with the given name from a list of Dino states.  If
no such dino exists, it returns nil.
*/
func (d DinoStates) FindDinoState(name string) *DinoState {
	for _, dino := range d {
		if dino.Dino.Name == name {
			return &dino
		}
	}
	return nil
}

/*
The DinoSensor type encodes the definition of a sensor.  Dinod only supports very simple
sensors at this stage, and the 'type' field is basically ignored.
*/
type DinoSensor struct {
	Name         string `json:"name"`
	FriendlyName string `json:"friendlyName"`
	Type         string `json:"type"`
	Pin          int    `json:"pin"`
}

/*
Valid checks whether a sensor definition is valid.  It returns a human-readable error, or
nil if the definition is valid.
*/
func (d DinoSensor) Valid() error {
	if d.Name == "" {
		return errors.New("sensor has no name")
	}

	if d.Type != "pulse" {
		return fmt.Errorf("sensor %v has an invalid type", d.Name)
	}

	return nil
}

/*
The DinoActuator type encodes the definition of a actuator.  Dinod only supports very simple
actuators at this stage, and the 'type' field is basically ignored.
*/
type DinoActuator struct {
	Name         string `json:"name"`
	FriendlyName string `json:"friendlyName"`
	Type         string `json:"type"`
	Pin          int    `json:"pin"`
}

/*
Valid checks whether an actuator definition is valid.  It returns a human-readable error,
or nil if the definition is valid.
*/
func (d DinoActuator) Valid() error {
	if d.Name == "" {
		return errors.New("actuator has no name")
	}

	if d.Type != "onoff" {
		return fmt.Errorf("actuator %v has an invalid type", d.Name)
	}

	return nil
}
