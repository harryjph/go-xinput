//+build !windows

package go_xinput

func getState(controllerIndex ControllerIndex) (*ControllerState, error) {
	return nil, UnsupportedOS
}
