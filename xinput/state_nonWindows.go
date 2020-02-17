//+build !windows

package xinput

func getState(controllerIndex ControllerIndex) (*ControllerState, error) {
	return nil, UnsupportedOS
}
