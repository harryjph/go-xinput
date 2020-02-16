//+build !windows

package go_xinput

func setState(controllerIndex ControllerIndex, leftMotorSpeed float32, rightMotorSpeed float32) error {
	return UnsupportedOS
}
