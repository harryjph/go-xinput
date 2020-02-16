//+build !windows

package go_xinput

func getControllerInfo(controllerIndex ControllerIndex) (*ControllerInfo, error) {
	return nil, UnsupportedOS
}
