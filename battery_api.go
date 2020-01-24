package go_xinput

type BatteryType uint8

const (
	Disconnected BatteryType = iota
	Wired
	Alkaline
	NiMH
	Unknown
)

type BatteryLevel uint8

const (
	Empty BatteryLevel = iota
	Low
	Medium
	Full
)

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
