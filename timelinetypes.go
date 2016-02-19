package main

type Timeline struct {
	FriendlyName string           `json:"friendlyName"`
	DinoName     string           `json:"dinoName"`
	Trigger      string           `json:"trigger"`
	Length       int              `json:"length"`
	TimePerStep  float64          `json:"timePerStep"`
	Timeline     map[string][]int `json:"timeline"`
}
