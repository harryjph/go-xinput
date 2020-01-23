package main

import (
	xinput "go-xinput"
	"log"
)

func main() {
	if xinput.LoadError != nil {
		log.Fatalln(xinput.LoadError)
	}

	log.Println("Controller:")
	batteryInfo, err := xinput.GetControllerBatteryInformation(xinput.Controller1)
	if err != nil {
		log.Fatalln(err)
	}
	printBatteryInfo(batteryInfo)

	log.Println("Headset:")
	batteryInfo, err = xinput.GetHeadsetBatteryInformation(xinput.Controller1)
	if err != nil {
		log.Fatalln(err)
	}
	printBatteryInfo(batteryInfo)
}

func printBatteryInfo(batteryInfo *xinput.BatteryInformation) {
	var batteryType string
	switch batteryInfo.BatteryType {
	case xinput.Disconnected: batteryType = "Disconnected"
	case xinput.Wired: batteryType = "Wired"
	case xinput.Alkaline: batteryType = "Alkaline"
	case xinput.NiMH: batteryType = "NiMH"
	case xinput.Unknown: batteryType = "Unknown"
	}

	var batteryLevel string
	switch batteryInfo.BatteryLevel {
	case xinput.Empty: batteryLevel = "Empty"
	case xinput.Low: batteryLevel = "Low"
	case xinput.Medium: batteryLevel = "Medium"
	case xinput.Full: batteryLevel = "Full"
	}

	log.Printf("Battery Type: %s, Battery Level: %s\n", batteryType, batteryLevel)
}
