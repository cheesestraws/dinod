package main

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTT_GPIO struct {
	client *mqtt.Client
}

// compile time check it complies with gpio interface
var _ = GPIO(&MQTT_GPIO{})

func (g *MQTT_GPIO) Init() {
	println("Initialising MQTT_GPIO...")
	// this is awful
	opts := mqtt.NewClientOptions().AddBroker("tcp://127.0.0.1:1883")
	g.client = mqtt.NewClient(opts)

	// connect
	token := g.client.Connect()
	token.Wait()
	if token.Error() != nil {
		fmt.Printf("giving up: %v\n", token.Error())
		return
	}
}

func (g *MQTT_GPIO) SetupInput(pin int, dinoName string, sensorName string) {
	// do nothing for the moment
	fmt.Printf("Pin %v is henceforth an output (%v, %v)\n", pin, dinoName, sensorName)
}

func (g *MQTT_GPIO) SetupOutput(pin int) {
	fmt.Printf("Pin %v is henceforth an output\n", pin)
}

func (g *MQTT_GPIO) SetPin(pin int, value int) {
	fmt.Printf("Pin %v => %v\n", pin, value)
	msg := fmt.Sprintf("%d", value)
	topic := fmt.Sprintf("dinod/output/%d", pin)

	token := g.client.Publish(topic, 0, false, msg)
	token.Wait()
	if token.Error() != nil {
		fmt.Printf("giving up: %v\n", token.Error())
		return
	}
}
