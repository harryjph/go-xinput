package main

import (
	"github.com/harry1453/go-xinput/xinput"
	"log"
)

func main() {
	if xinput.LoadError != nil {
		log.Fatalln(xinput.LoadError)
	}

	batteryInfo, err := xinput.GetControllerBatteryInformation(xinput.Controller1)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Controller: Battery Type: %s, Battery Level: %s\n", batteryInfo.BatteryType, batteryInfo.BatteryLevel)

	batteryInfo, err = xinput.GetHeadsetBatteryInformation(xinput.Controller1)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Headset: Battery Type: %s, Battery Level: %s\n", batteryInfo.BatteryType, batteryInfo.BatteryLevel)
}
