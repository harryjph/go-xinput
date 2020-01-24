package main

import (
	xInput "go-xinput"
	"log"
)

func main() {
	if xInput.LoadError != nil {
		log.Fatalln(xInput.LoadError)
	}

	log.Println("Controller:")
	batteryInfo, err := xInput.GetControllerBatteryInformation(xInput.Controller1)
	if err != nil {
		log.Fatalln(err)
	}
	printBatteryInfo(batteryInfo)

	log.Println("Headset:")
	batteryInfo, err = xInput.GetHeadsetBatteryInformation(xInput.Controller1)
	if err != nil {
		log.Fatalln(err)
	}
	printBatteryInfo(batteryInfo)
}

func printBatteryInfo(batteryInfo *xInput.BatteryInformation) {
	var batteryType string
	switch batteryInfo.BatteryType {
	case xInput.Disconnected:
		batteryType = "Disconnected"
	case xInput.Wired:
		batteryType = "Wired"
	case xInput.Alkaline:
		batteryType = "Alkaline"
	case xInput.NiMH:
		batteryType = "NiMH"
	case xInput.Unknown:
		batteryType = "Unknown"
	}

	var batteryLevel string
	switch batteryInfo.BatteryLevel {
	case xInput.Empty:
		batteryLevel = "Empty"
	case xInput.Low:
		batteryLevel = "Low"
	case xInput.Medium:
		batteryLevel = "Medium"
	case xInput.Full:
		batteryLevel = "Full"
	}

	log.Printf("Battery Type: %s, Battery Level: %s\n", batteryType, batteryLevel)
}
