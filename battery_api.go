package go_xinput

import "fmt"

type BatteryType uint8

const (
	Disconnected BatteryType = iota
	Wired
	Alkaline
	NiMH
	Unknown
)

func (batteryType BatteryType) String() string {
	switch batteryType {
	case Disconnected:
		return "Disconnected"
	case Wired:
		return "Wired"
	case Alkaline:
		return "Alkaline"
	case NiMH:
		return "NiMH"
	case Unknown:
		return "Unknown"
	default:
		return fmt.Sprintf("Unrecognized BatteryType: %02x", uint8(batteryType))
	}
}

type BatteryLevel uint8

const (
	Empty BatteryLevel = iota
	Low
	Medium
	Full
)

func (batteryLevel BatteryLevel) String() string {
	switch batteryLevel {
	case Empty:
		return "Empty"
	case Low:
		return "Low"
	case Medium:
		return "Medium"
	case Full:
		return "Full"
	default:
		return fmt.Sprintf("Unrecognized BatteryLevel: %02x", uint8(batteryLevel))
	}
}

type BatteryInformation struct {
	BatteryType  BatteryType
	BatteryLevel BatteryLevel
}

func GetControllerBatteryInformation(controllerIndex ControllerIndex) (*BatteryInformation, error) {
	return getBatteryInformation(controllerIndex, xInputBatteryDevtypeGamepad)
}

func GetHeadsetBatteryInformation(controllerIndex ControllerIndex) (*BatteryInformation, error) {
	return getBatteryInformation(controllerIndex, xInputBatteryDevtypeHeadset)
}
