package main

import (
	xinput "go-xinput"
	"log"
	"time"
)

func main() {
	if xinput.LoadError != nil {
		log.Fatalln(xinput.LoadError)
	}

	log.Println("Setting vibration to 25%...")
	if err := xinput.SetVibration(xinput.Controller1, 0.25, 0.25); err != nil {
		log.Fatalln(err)
	}

	time.Sleep(time.Second)

	log.Println("Setting vibration to 50%...")
	if err := xinput.SetVibration(xinput.Controller1, 0.50, 0.50); err != nil {
		log.Fatalln(err)
	}

	time.Sleep(time.Second)

	log.Println("Setting vibration to 75%...")
	if err := xinput.SetVibration(xinput.Controller1, 0.75, 0.75); err != nil {
		log.Fatalln(err)
	}

	time.Sleep(time.Second)

	log.Println("Setting vibration to 100%...")
	if err := xinput.SetVibration(xinput.Controller1, 1.00, 1.00); err != nil {
		log.Fatalln(err)
	}

	time.Sleep(time.Second)

	log.Println("Setting vibration to 0%...")
	if err := xinput.SetVibration(xinput.Controller1, 0, 0); err != nil {
		log.Fatalln(err)
	}
}
