//+build windows

package go_xinput

import (
	"syscall"
	"unsafe"
)

const (
	xInputDevSubTypeUnknown         = 0x00
	xInputDevSubTypeGamepad         = 0x01
	xInputDevSubTypeWheel           = 0x02
	xInputDevSubTypeArcadeStick     = 0x03
	xInputDevSubTypeFlightStick     = 0x04
	xInputDevSubTypeDancePad        = 0x05
	xInputDevSubTypeGuitar          = 0x06
	xInputDevSubTypeGuitarAlternate = 0x07
	xInputDevSubTypeDrumKit         = 0x08
	xInputDevSubTypeGuitarBass      = 0x0B
	xInputDevSubTypeArcadePad       = 0x13

	xInputCapsVoiceSupported = 0x0004
	xInputCapsFfbSupported   = 0x0001
	xInputCapsWireless       = 0x0002
	xInputCapsPmdSupported   = 0x0008
	xInputCapsNoNavigation   = 0x0010
)

func devSubTypeToControllerType(subtype uint8) ControllerType {
	switch subtype {
	case xInputDevSubTypeUnknown:
		return UnknownType
	case xInputDevSubTypeGamepad:
		return Gamepad
	case xInputDevSubTypeWheel:
		return Wheel
	case xInputDevSubTypeArcadeStick:
		return ArcadeStick
	case xInputDevSubTypeFlightStick:
		return FlightStick
	case xInputDevSubTypeDancePad:
		return DancePad
	case xInputDevSubTypeGuitar:
		return Guitar
	case xInputDevSubTypeGuitarAlternate:
		return GuitarAlternative
	case xInputDevSubTypeDrumKit:
		return DrumKit
	case xInputDevSubTypeGuitarBass:
		return GuitarBass
	case xInputDevSubTypeArcadePad:
		return ArcadePad
	default:
		return UnknownType
	}
}

type xInputCapabilities struct {
	controllerType uint8
	subtype        uint8
	flags          uint16
	state          xInputState
	vibration      xInputVibration
}

func (capabilities *xInputCapabilities) toControllerInfo() *ControllerInfo {
	return &ControllerInfo{
		Subtype: devSubTypeToControllerType(capabilities.subtype),
		Features: ControllerFeatures{
			VoiceSupported:         capabilities.flags&xInputCapsVoiceSupported != 0,
			VibrationSupported:     capabilities.flags&xInputCapsFfbSupported != 0,
			Wireless:               capabilities.flags&xInputCapsWireless != 0,
			PluginModulesSupported: capabilities.flags&xInputCapsPmdSupported != 0,
			HasNavigationButtons:   capabilities.flags&xInputCapsNoNavigation == 0,
		},
		State:     *capabilities.state.toControllerState(),
		Vibration: capabilities.vibration.toControllerVibration(),
	}
}

func getControllerInfo(controllerIndex ControllerIndex) (*ControllerInfo, error) {
	if xInputGetCapabilities == nil {
		return nil, FunctionNotAvailable
	}

	userIndex, err := controllerIndex.toUserIndex()
	if err != nil {
		return nil, err
	}

	var capabilities xInputCapabilities

	r, _, _ := xInputGetCapabilities.Call(userIndex, 0, uintptr(unsafe.Pointer(&capabilities)))
	if r == 0 {
		return capabilities.toControllerInfo(), nil
	} else {
		return nil, syscall.Errno(r)
	}
}
