package go_xinput

import (
	"syscall"
	"unsafe"
)

const (
	xInputGamepadDpadUp        = 0x0001
	xInputGamepadDpadDown      = 0x0002
	xInputGamepadDpadLeft      = 0x0004
	xInputGamepadDpadRight     = 0x0008
	xInputGamepadStart         = 0x0010
	xInputGamepadBack          = 0x0020
	xInputGamepadLeftThumb     = 0x0040
	xInputGamepadRightThumb    = 0x0080
	xInputGamepadLeftShoulder  = 0x0100
	xInputGamepadRightShoulder = 0x0200
	xInputGamepadA             = 0x1000
	xInputGamepadB             = 0x2000
	xInputGamepadX             = 0x4000
	xInputGamepadY             = 0x8000
)

type xInputGamepad struct {
	packetNumber uint32
	state        xInputState
}

type xInputState struct {
	buttons      uint16
	leftTrigger  uint8
	rightTrigger uint8
	thumbLX      int16
	thumbLY      int16
	thumbRX      int16
	thumbRY      int16
}

// Deadzones
const (
	triggerDeadzone = 30
)

func uint8ToFloat(uint uint8) float32 {
	if uint < triggerDeadzone {
		return 0
	}
	return float32(uint) / float32(0xFF)
}

func int16ToFloat(int int16) float32 {
	// TODO deadzone
	float := float32(int) / float32(0x7FFF)
	if float < -1 {
		return -1
	}
	if float > 1 {
		return 1
	}
	return float
}

func (state *xInputState) toControllerState() *ControllerState {
	return &ControllerState{
		Buttons: Buttons{
			A:             state.buttons&xInputGamepadA != 0,
			B:             state.buttons&xInputGamepadB != 0,
			X:             state.buttons&xInputGamepadX != 0,
			Y:             state.buttons&xInputGamepadY != 0,
			DpadUp:        state.buttons&xInputGamepadDpadUp != 0,
			DpadDown:      state.buttons&xInputGamepadDpadDown != 0,
			DpadLeft:      state.buttons&xInputGamepadDpadLeft != 0,
			DpadRight:     state.buttons&xInputGamepadDpadRight != 0,
			LeftShoulder:  state.buttons&xInputGamepadLeftShoulder != 0,
			RightShoulder: state.buttons&xInputGamepadRightShoulder != 0,
			LeftJoystick:  state.buttons&xInputGamepadLeftThumb != 0,
			RightJoystick: state.buttons&xInputGamepadRightThumb != 0,
			Start:         state.buttons&xInputGamepadStart != 0,
			Back:          state.buttons&xInputGamepadBack != 0,
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

	var gamepad xInputGamepad
	r, _, _ := xInputGetState.Call(userIndex, uintptr(unsafe.Pointer(&gamepad)))
	if r == 0 {
		return gamepad.state.toControllerState(), nil
	} else {
		return nil, syscall.Errno(r)
	}
}
