//+build windows

package xinput

import (
	"fmt"
	"syscall"
	"unsafe"
)

type xInputVibration struct {
	leftMotorSpeed  uint16
	rightMotorSpeed uint16
}

func (vibration *xInputVibration) toControllerVibration() ControllerVibration {
	return ControllerVibration{
		LowFrequencyLevel:  uint16ToFloat(vibration.leftMotorSpeed),
		HighFrequencyLevel: uint16ToFloat(vibration.rightMotorSpeed),
	}
}

func uint16ToFloat(uint uint16) float32 {
	return float32(uint) / float32(uint16(0xFFFF))
}

func floatToUint16(float float32) (uint16, error) {
	if float < 0 {
		return 0, fmt.Errorf("motor speed value must be 0-1, was %.3f", float)
	}
	if float > 1 {
		return 0, fmt.Errorf("motor speed value must be 0-1, was %.3f", float)
	}
	return uint16(float * 0xFFFF), nil
}

func setState(controllerIndex ControllerIndex, leftMotorSpeed float32, rightMotorSpeed float32) error {
	if xInputSetState == nil {
		return FunctionNotAvailable
	}

	userIndex, err := controllerIndex.toUserIndex()
	if err != nil {
		return err
	}

	leftMotorSpeedInt, err := floatToUint16(leftMotorSpeed)
	if err != nil {
		return err
	}
	rightMotorSpeedInt, err := floatToUint16(rightMotorSpeed)
	if err != nil {
		return err
	}

	r, _, _ := xInputSetState.Call(userIndex, uintptr(unsafe.Pointer(&xInputVibration{
		leftMotorSpeed:  leftMotorSpeedInt,
		rightMotorSpeed: rightMotorSpeedInt,
	})))

	if r == 0 {
		return nil
	} else {
		return syscall.Errno(r)
	}
}
