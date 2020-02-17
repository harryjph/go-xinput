//+build !windows

package xinput

func setState(controllerIndex ControllerIndex, leftMotorSpeed float32, rightMotorSpeed float32) error {
	return UnsupportedOS
}
