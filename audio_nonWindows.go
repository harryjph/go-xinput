//+build !windows

package go_xinput

func getAudioDeviceIds(controllerIndex ControllerIndex) (inputDeviceId, outputDeviceId []byte, err error) {
	return nil, nil, UnsupportedOS
}
