//+build !windows

package xinput

func getControllerInfo(controllerIndex ControllerIndex) (*ControllerInfo, error) {
	return nil, UnsupportedOS
}
