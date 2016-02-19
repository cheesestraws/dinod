package main

type Dino struct {
	Name         string         `json:"name"`
	FriendlyName string         `json:"friendlyName"`
	ImageURI     string         `json:"imageURI"`
	Sensors      []DinoSensor   `json:"sensors"`
	Actuators    []DinoActuator `json:"actuators"`
}

func (d Dino) Valid() bool {
	if d.Name == "" {
		return false
	}

	for _, sensor := range d.Sensors {
		if !sensor.Valid() {
			return false
		}
	}

	for _, actuator := range d.Actuators {
		if !actuator.Valid() {
			return false
		}
	}

	return true
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

func (d DinoSensor) Valid() bool {
	if d.Name == "" {
		return false
	}

	if d.Type != "pulse" {
		return false
	}

	return true
}

type DinoActuator struct {
	Name         string `json:"name"`
	FriendlyName string `json:"friendlyName"`
	Type         string `json:"type"`
	Pin          int    `json:"pin"`
}

func (d DinoActuator) Valid() bool {
	if d.Name == "" {
		return false
	}

	if d.Type != "onoff" {
		return false
	}

	return true
}
