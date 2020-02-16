//+build windows

package go_xinput

import (
	"syscall"
	"unsafe"
)

// TODO This currently always returns ERROR_DEVICE_NOT_CONNECTED.
func getAudioDeviceIds(controllerIndex ControllerIndex) (inputDeviceId, outputDeviceId []byte, err error) {
	if xInputGetAudioDeviceIds == nil {
		return nil, nil, FunctionNotAvailable
	}
	userIndex, err := controllerIndex.toUserIndex()
	if err != nil {
		return nil, nil, err
	}
	size := 256
	inputDeviceId = make([]byte, size)
	outputDeviceId = make([]byte, size)
	r, _, _ := xInputGetAudioDeviceIds.Call(userIndex, uintptr(unsafe.Pointer(&outputDeviceId[0])), uintptr(unsafe.Pointer(&size)), uintptr(unsafe.Pointer(&inputDeviceId[0])), uintptr(unsafe.Pointer(&size)))
	if r == 0 {
		return
	} else {
		return nil, nil, syscall.Errno(r)
	}
}
