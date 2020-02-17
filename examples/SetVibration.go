package main

import (
	"github.com/harry1453/go-xinput/xinput"
	"log"
	"time"
)

func main() {
	if xinput.LoadError != nil {
		log.Fatalln(xinput.LoadError)
	}

	log.Println("Setting vibration to 25%...")
	if err := xinput.SetVibration(xinput.Controller1, xinput.ControllerVibration{LowFrequencyLevel: 0.25, HighFrequencyLevel: 0.25}); err != nil {
		log.Fatalln(err)
	}

	time.Sleep(time.Second)

	log.Println("Setting vibration to 50%...")
	if err := xinput.SetVibration(xinput.Controller1, xinput.ControllerVibration{LowFrequencyLevel: 0.50, HighFrequencyLevel: 0.50}); err != nil {
		log.Fatalln(err)
	}

	time.Sleep(time.Second)

	log.Println("Setting vibration to 75%...")
	if err := xinput.SetVibration(xinput.Controller1, xinput.ControllerVibration{LowFrequencyLevel: 0.75, HighFrequencyLevel: 0.75}); err != nil {
		log.Fatalln(err)
	}

	time.Sleep(time.Second)

	log.Println("Setting vibration to 100%...")
	if err := xinput.SetVibration(xinput.Controller1, xinput.ControllerVibration{LowFrequencyLevel: 1, HighFrequencyLevel: 1}); err != nil {
		log.Fatalln(err)
	}

	time.Sleep(time.Second)

	log.Println("Setting vibration to 0%...")
	if err := xinput.SetVibration(xinput.Controller1, xinput.ControllerVibration{LowFrequencyLevel: 0, HighFrequencyLevel: 0}); err != nil {
		log.Fatalln(err)
	}
}
