package main

type GPIO interface {
	Init()
	SetupInput(pin int, dinoName string, sensorName string)
	SetupOutput(pin int)
	SetPin(pin int, value int)
}

func SetupGPIOForDinos(dinos Dinos, g GPIO) {
	// iterate over each dino
	for _, dino := range dinos {
		// Sensors are all input.
		for _, sensor := range dino.Sensors {
			g.SetupInput(sensor.Pin)
		}
		for _, actuator := range dino.Actuators {
			g.SetupOutput(actuator.Pin)
		}
	}
}
