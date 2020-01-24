package go_xinput

import (
	"syscall"
	"unsafe"
)

const (
	xinputGamepadDpadUp        = 0x0001
	xinputGamepadDpadDown      = 0x0002
	xinputGamepadDpadLeft      = 0x0004
	xinputGamepadDpadRight     = 0x0008
	xinputGamepadStart         = 0x0010
	xinputGamepadBack          = 0x0020
	xinputGamepadLeftThumb     = 0x0040
	xinputGamepadRightThumb    = 0x0080
	xinputGamepadLeftShoulder  = 0x0100
	xinputGamepadRightShoulder = 0x0200
	xinputGamepadA             = 0x1000
	xinputGamepadB             = 0x2000
	xinputGamepadX             = 0x4000
	xinputGamepadY             = 0x8000
)

type xinputGamepad struct {
	packetNumber uint32
	state        xinputState
}

type xinputState struct {
	buttons      uint16
	leftTrigger  uint8
	rightTrigger uint8
	thumbLX      int16
	thumbLY      int16
	thumbRX      int16
	thumbRY      int16
}

func uint8ToFloat(uint uint8) float32 {
	return float32(uint) / float32(0xFF)
}

func int16ToFloat(int int16) float32 {
	float := float32(int) / float32(0x7FFF)
	if float < -1 {
		return -1
	}
	if float > 1 {
		return 1
	}
	return float
}

func (state *xinputState) toControllerState() *ControllerState {
	return &ControllerState{
		Buttons: Buttons{
			A:             state.buttons&xinputGamepadA != 0,
			B:             state.buttons&xinputGamepadB != 0,
			X:             state.buttons&xinputGamepadX != 0,
			Y:             state.buttons&xinputGamepadY != 0,
			DpadUp:        state.buttons&xinputGamepadDpadUp != 0,
			DpadDown:      state.buttons&xinputGamepadDpadDown != 0,
			DpadLeft:      state.buttons&xinputGamepadDpadLeft != 0,
			DpadRight:     state.buttons&xinputGamepadDpadRight != 0,
			LeftShoulder:  state.buttons&xinputGamepadLeftShoulder != 0,
			RightShoulder: state.buttons&xinputGamepadRightShoulder != 0,
			LeftJoystick:  state.buttons&xinputGamepadLeftThumb != 0,
			RightJoystick: state.buttons&xinputGamepadRightThumb != 0,
			Start:         state.buttons&xinputGamepadStart != 0,
			Back:          state.buttons&xinputGamepadBack != 0,
		},
		LeftJoystick: Joystick{
			X: int16ToFloat(state.thumbLX),
			Y: int16ToFloat(state.thumbLY),
		},
		RightJoystick: Joystick{
			X: int16ToFloat(state.thumbRX),
			Y: int16ToFloat(state.thumbRY),
		},
		LeftTrigger:  uint8ToFloat(state.leftTrigger),
		RightTrigger: uint8ToFloat(state.rightTrigger),
	}
}

func getState(controllerIndex ControllerIndex) (*ControllerState, error) {
	if xInputGetState == nil {
		return nil, FunctionNotAvailable
	}

	userIndex, err := controllerIndex.toUserIndex()
	if err != nil {
		return nil, err
	}

	var gamepad xinputGamepad
	r, _, _ := xInputGetState.Call(userIndex, uintptr(unsafe.Pointer(&gamepad)))
	if r == 0 {
		return gamepad.state.toControllerState(), nil
	} else {
		return nil, syscall.Errno(r)
	}
}
