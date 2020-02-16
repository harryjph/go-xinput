//+build !windows

package go_xinput

func getControllerBatteryInformation(controllerIndex ControllerIndex) (*BatteryInformation, error) {
	return nil, UnsupportedOS
}

func getHeadsetBatteryInformation(controllerIndex ControllerIndex) (*BatteryInformation, error) {
	return nil, UnsupportedOS
}
